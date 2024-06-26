// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package gitserver

import (
	"context"
	"sync"
	"time"

	api "github.com/sourcegraph/sourcegraph/internal/api"
)

// MockRepositoryServiceClient is a mock implementation of the
// RepositoryServiceClient interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/repo-updater/internal/gitserver)
// used for unit testing.
type MockRepositoryServiceClient struct {
	// DeleteRepositoryFunc is an instance of a mock function object
	// controlling the behavior of the method DeleteRepository.
	DeleteRepositoryFunc *RepositoryServiceClientDeleteRepositoryFunc
	// FetchRepositoryFunc is an instance of a mock function object
	// controlling the behavior of the method FetchRepository.
	FetchRepositoryFunc *RepositoryServiceClientFetchRepositoryFunc
}

// NewMockRepositoryServiceClient creates a new mock of the
// RepositoryServiceClient interface. All methods return zero values for all
// results, unless overwritten.
func NewMockRepositoryServiceClient() *MockRepositoryServiceClient {
	return &MockRepositoryServiceClient{
		DeleteRepositoryFunc: &RepositoryServiceClientDeleteRepositoryFunc{
			defaultHook: func(context.Context, api.RepoName) (r0 error) {
				return
			},
		},
		FetchRepositoryFunc: &RepositoryServiceClientFetchRepositoryFunc{
			defaultHook: func(context.Context, api.RepoName) (r0 time.Time, r1 time.Time, r2 error) {
				return
			},
		},
	}
}

// NewStrictMockRepositoryServiceClient creates a new mock of the
// RepositoryServiceClient interface. All methods panic on invocation,
// unless overwritten.
func NewStrictMockRepositoryServiceClient() *MockRepositoryServiceClient {
	return &MockRepositoryServiceClient{
		DeleteRepositoryFunc: &RepositoryServiceClientDeleteRepositoryFunc{
			defaultHook: func(context.Context, api.RepoName) error {
				panic("unexpected invocation of MockRepositoryServiceClient.DeleteRepository")
			},
		},
		FetchRepositoryFunc: &RepositoryServiceClientFetchRepositoryFunc{
			defaultHook: func(context.Context, api.RepoName) (time.Time, time.Time, error) {
				panic("unexpected invocation of MockRepositoryServiceClient.FetchRepository")
			},
		},
	}
}

// NewMockRepositoryServiceClientFrom creates a new mock of the
// MockRepositoryServiceClient interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockRepositoryServiceClientFrom(i RepositoryServiceClient) *MockRepositoryServiceClient {
	return &MockRepositoryServiceClient{
		DeleteRepositoryFunc: &RepositoryServiceClientDeleteRepositoryFunc{
			defaultHook: i.DeleteRepository,
		},
		FetchRepositoryFunc: &RepositoryServiceClientFetchRepositoryFunc{
			defaultHook: i.FetchRepository,
		},
	}
}

// RepositoryServiceClientDeleteRepositoryFunc describes the behavior when
// the DeleteRepository method of the parent MockRepositoryServiceClient
// instance is invoked.
type RepositoryServiceClientDeleteRepositoryFunc struct {
	defaultHook func(context.Context, api.RepoName) error
	hooks       []func(context.Context, api.RepoName) error
	history     []RepositoryServiceClientDeleteRepositoryFuncCall
	mutex       sync.Mutex
}

// DeleteRepository delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockRepositoryServiceClient) DeleteRepository(v0 context.Context, v1 api.RepoName) error {
	r0 := m.DeleteRepositoryFunc.nextHook()(v0, v1)
	m.DeleteRepositoryFunc.appendCall(RepositoryServiceClientDeleteRepositoryFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the DeleteRepository
// method of the parent MockRepositoryServiceClient instance is invoked and
// the hook queue is empty.
func (f *RepositoryServiceClientDeleteRepositoryFunc) SetDefaultHook(hook func(context.Context, api.RepoName) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteRepository method of the parent MockRepositoryServiceClient
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *RepositoryServiceClientDeleteRepositoryFunc) PushHook(hook func(context.Context, api.RepoName) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryServiceClientDeleteRepositoryFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, api.RepoName) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryServiceClientDeleteRepositoryFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, api.RepoName) error {
		return r0
	})
}

func (f *RepositoryServiceClientDeleteRepositoryFunc) nextHook() func(context.Context, api.RepoName) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryServiceClientDeleteRepositoryFunc) appendCall(r0 RepositoryServiceClientDeleteRepositoryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// RepositoryServiceClientDeleteRepositoryFuncCall objects describing the
// invocations of this function.
func (f *RepositoryServiceClientDeleteRepositoryFunc) History() []RepositoryServiceClientDeleteRepositoryFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryServiceClientDeleteRepositoryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryServiceClientDeleteRepositoryFuncCall is an object that
// describes an invocation of method DeleteRepository on an instance of
// MockRepositoryServiceClient.
type RepositoryServiceClientDeleteRepositoryFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 api.RepoName
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryServiceClientDeleteRepositoryFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryServiceClientDeleteRepositoryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// RepositoryServiceClientFetchRepositoryFunc describes the behavior when
// the FetchRepository method of the parent MockRepositoryServiceClient
// instance is invoked.
type RepositoryServiceClientFetchRepositoryFunc struct {
	defaultHook func(context.Context, api.RepoName) (time.Time, time.Time, error)
	hooks       []func(context.Context, api.RepoName) (time.Time, time.Time, error)
	history     []RepositoryServiceClientFetchRepositoryFuncCall
	mutex       sync.Mutex
}

// FetchRepository delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockRepositoryServiceClient) FetchRepository(v0 context.Context, v1 api.RepoName) (time.Time, time.Time, error) {
	r0, r1, r2 := m.FetchRepositoryFunc.nextHook()(v0, v1)
	m.FetchRepositoryFunc.appendCall(RepositoryServiceClientFetchRepositoryFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the FetchRepository
// method of the parent MockRepositoryServiceClient instance is invoked and
// the hook queue is empty.
func (f *RepositoryServiceClientFetchRepositoryFunc) SetDefaultHook(hook func(context.Context, api.RepoName) (time.Time, time.Time, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// FetchRepository method of the parent MockRepositoryServiceClient instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *RepositoryServiceClientFetchRepositoryFunc) PushHook(hook func(context.Context, api.RepoName) (time.Time, time.Time, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryServiceClientFetchRepositoryFunc) SetDefaultReturn(r0 time.Time, r1 time.Time, r2 error) {
	f.SetDefaultHook(func(context.Context, api.RepoName) (time.Time, time.Time, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryServiceClientFetchRepositoryFunc) PushReturn(r0 time.Time, r1 time.Time, r2 error) {
	f.PushHook(func(context.Context, api.RepoName) (time.Time, time.Time, error) {
		return r0, r1, r2
	})
}

func (f *RepositoryServiceClientFetchRepositoryFunc) nextHook() func(context.Context, api.RepoName) (time.Time, time.Time, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryServiceClientFetchRepositoryFunc) appendCall(r0 RepositoryServiceClientFetchRepositoryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// RepositoryServiceClientFetchRepositoryFuncCall objects describing the
// invocations of this function.
func (f *RepositoryServiceClientFetchRepositoryFunc) History() []RepositoryServiceClientFetchRepositoryFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryServiceClientFetchRepositoryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryServiceClientFetchRepositoryFuncCall is an object that
// describes an invocation of method FetchRepository on an instance of
// MockRepositoryServiceClient.
type RepositoryServiceClientFetchRepositoryFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 api.RepoName
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 time.Time
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 time.Time
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryServiceClientFetchRepositoryFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryServiceClientFetchRepositoryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}
