package observation

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"go.opentelemetry.io/otel/attribute"
)

// commonAcronyms includes acronyms that malform the expected output of kebabCase
// due to unexpected adjacent upper-case letters. Add items to this list to stop
// kebabCase from transforming `FromLSIF` into `from-l-s-i-f`.
var commonAcronyms = []string{
	"API",
	"ID",
	"LSIF",
}

// acronymsReplacer is a string replacer that normalizes the acronyms above. For
// example, `API` will be transformed into `Api` so that it appears as one word.
var acronymsReplacer *strings.Replacer

func init() {
	var pairs []string
	for _, acronym := range commonAcronyms {
		pairs = append(pairs, acronym, fmt.Sprintf("%c%s", acronym[0], strings.ToLower(acronym[1:])))
	}

	acronymsReplacer = strings.NewReplacer(pairs...)
}

// kebab transforms a string into lower-kebab-case.
func kebabCase(s string) string {
	// Normalize all acronyms before looking at character transitions
	s = acronymsReplacer.Replace(s)

	buf := bytes.NewBufferString("")
	for i, c := range s {
		// If we've seen a letter and we're going lower -> upper, add a skewer
		if i > 0 && unicode.IsLower(rune(s[i-1])) && unicode.IsUpper(c) {
			buf.WriteRune('-')
		}

		buf.WriteRune(unicode.ToLower(c))
	}

	return buf.String()
}

// mergeLabels flattens slices of slices of strings.
func mergeLabels(groups ...[]string) []string {
	size := 0
	for _, group := range groups {
		size += len(group)
	}

	labels := make([]string, 0, size)
	for _, group := range groups {
		labels = append(labels, group...)
	}

	return labels
}

// mergeAttrs flattens slices of slices of log fields.
func mergeAttrs(groups ...[]attribute.KeyValue) []attribute.KeyValue {
	size := 0
	for _, group := range groups {
		size += len(group)
	}

	attrs := make([]attribute.KeyValue, 0, size)
	for _, group := range groups {
		attrs = append(attrs, group...)
	}

	return attrs
}
