package internal

import (
	"context"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sourcegraph/log"
	"go.opentelemetry.io/otel/attribute"

	"github.com/sourcegraph/sourcegraph/cmd/gitserver/internal/common"
	"github.com/sourcegraph/sourcegraph/cmd/gitserver/internal/git/gitcli"
	"github.com/sourcegraph/sourcegraph/cmd/gitserver/internal/search"
	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/gitserver/protocol"
	"github.com/sourcegraph/sourcegraph/internal/honey"
	"github.com/sourcegraph/sourcegraph/internal/trace"
)

var (
	searchRunning = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "src_gitserver_search_running",
		Help: "number of gitserver.Search running concurrently.",
	})
	searchDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "src_gitserver_search_duration_seconds",
		Help:    "gitserver.Search duration in seconds.",
		Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 30},
	}, []string{"error"})
	searchLatency = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "src_gitserver_search_latency_seconds",
		Help:    "gitserver.Search latency (time until first result is sent) in seconds.",
		Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 30},
	})
)

func searchWithObservability(ctx context.Context, logger log.Logger, repoDir common.GitDir, tr trace.Trace, args *protocol.SearchRequest, onMatch func(*protocol.CommitMatch) error) (limitHit bool, err error) {
	searchStart := time.Now()

	searchRunning.Inc()
	defer searchRunning.Dec()

	tr.SetAttributes(
		args.Repo.Attr(),
		attribute.Bool("include_diff", args.IncludeDiff),
		attribute.String("query", args.Query.String()),
		attribute.Int("limit", args.Limit),
		attribute.Bool("include_modified_files", args.IncludeModifiedFiles),
	)
	defer func() {
		tr.AddEvent("done", attribute.Bool("limit_hit", limitHit))
		tr.SetError(err)
		searchDuration.
			WithLabelValues(strconv.FormatBool(err != nil)).
			Observe(time.Since(searchStart).Seconds())

		if honey.Enabled() || traceLogs {
			act := actor.FromContext(ctx)
			ev := honey.NewEvent("gitserver-search")
			ev.SetSampleRate(gitcli.HoneySampleRate("", act))
			ev.AddField("repo", args.Repo)
			ev.AddField("revisions", args.Revisions)
			ev.AddField("include_diff", args.IncludeDiff)
			ev.AddField("include_modified_files", args.IncludeModifiedFiles)
			ev.AddField("actor", act.UIDString())
			ev.AddField("query", args.Query.String())
			ev.AddField("limit", args.Limit)
			ev.AddField("duration_ms", time.Since(searchStart).Milliseconds())
			if err != nil {
				ev.AddField("error", err.Error())
			}
			if traceID := trace.ID(ctx); traceID != "" {
				ev.AddField("traceID", traceID)
				ev.AddField("trace", trace.URL(traceID, conf.DefaultClient()))
			}
			if honey.Enabled() {
				_ = ev.Send()
			}
			if traceLogs {
				logger.Debug("TRACE gitserver search", log.Object("ev.Fields", mapToLoggerField(ev.Fields())...))
			}
		}
	}()

	observeLatency := sync.OnceFunc(func() {
		searchLatency.Observe(time.Since(searchStart).Seconds())
	})

	onMatchWithLatency := func(cm *protocol.CommitMatch) error {
		observeLatency()
		return onMatch(cm)
	}

	return doSearch(ctx, logger, repoDir, args, onMatchWithLatency)
}

// doSearch handles the core logic of the search. It is passed a matchesBuf so it doesn't need to
// concern itself with event types, and all instrumentation is handled in the calling function.
func doSearch(ctx context.Context, logger log.Logger, repoDir common.GitDir, args *protocol.SearchRequest, onMatch func(*protocol.CommitMatch) error) (limitHit bool, err error) {
	if args.Limit == 0 {
		args.Limit = math.MaxInt32
	}

	mt, err := search.ToMatchTree(args.Query)
	if err != nil {
		return false, err
	}

	// Ensure that we populate ModifiedFiles when we have a DiffModifiesFile filter.
	// --name-status is not zero cost, so we don't do it on every search.
	hasDiffModifiesFile := false
	search.Visit(mt, func(mt search.MatchTree) {
		switch mt.(type) {
		case *search.DiffModifiesFile:
			hasDiffModifiesFile = true
		}
	})

	// Create a recvCallback that detects whether we've hit a limit
	// and stops sending when we have.
	var sentCount atomic.Int64
	var hitLimit atomic.Bool
	limitedOnMatch := func(match *protocol.CommitMatch) {
		// Avoid sending if we've already hit the limit
		if int(sentCount.Load()) >= args.Limit {
			hitLimit.Store(true)
			return
		}

		sentCount.Add(int64(matchCount(match)))
		_ = onMatch(match)
	}

	searcher := &search.CommitSearcher{
		Logger:               logger,
		RepoName:             args.Repo,
		RepoDir:              string(repoDir),
		Revisions:            args.Revisions,
		Query:                mt,
		IncludeDiff:          args.IncludeDiff,
		IncludeModifiedFiles: args.IncludeModifiedFiles || hasDiffModifiesFile,
	}

	return hitLimit.Load(), searcher.Search(ctx, limitedOnMatch)
}

// matchCount returns either:
// 1) the number of diff matches if there are any
// 2) the number of messsage matches if there are any
// 3) one, to represent matching the commit, but nothing inside it
func matchCount(cm *protocol.CommitMatch) int {
	if len(cm.Diff.MatchedRanges) > 0 {
		return len(cm.Diff.MatchedRanges)
	}
	if len(cm.Message.MatchedRanges) > 0 {
		return len(cm.Message.MatchedRanges)
	}
	return 1
}
