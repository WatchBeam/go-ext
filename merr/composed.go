package merr

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

// StackTracer is an error type that returns stacktraces from pkg/errors.
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// ComposedError is an error that encapsulates and prints multiple errors.
type ComposedError []error

var _ error = ComposedError{}

// Compose creates a new ComposedError from the list of individual errors.
func Compose(errs ...error) *ComposedError { list := ComposedError(errs); return &list }

// Add inserts a new error into the composed list. If err is nil, it's a noop.
func (c *ComposedError) Add(errs ...error) *ComposedError {
	for _, err := range errs {
		if err != nil {
			*c = append(*c, err)
		}
	}

	return c
}

// Empty returns true if the composed error has no nested errors.
func (c ComposedError) Empty() bool { return len(c) == 0 }

// Error implements error.Error
func (c ComposedError) Error() string {
	var msg string
	if len(c) == 0 {
		return "go-ext/composed: no errors recorded"
	}

	for i, err := range c {
		if i > 0 {
			msg += "\n"
		}

		msg += "error #" + strconv.Itoa(i) + ": "
		if st, ok := err.(StackTracer); ok {
			msg += fmt.Sprintf("%+v", st)
		} else {
			msg += err.Error()
		}
	}

	return msg
}

// ConcurrentComposedError is a version of ComposedError safe for concurrent usage.
type ConcurrentComposedError struct {
	ComposedError
	mu sync.Mutex
}

// Compose creates a new ConcurrentComposedError from the list of individual errors.
func ComposeConcurrent(errs ...error) *ConcurrentComposedError {
	return &ConcurrentComposedError{ComposedError: *Compose(errs...)}
}

// Add inserts a new error into the composed list. If err is nil, it's a noop.
func (c *ConcurrentComposedError) Add(errs ...error) *ConcurrentComposedError {
	c.mu.Lock()
	c.ComposedError.Add(errs...)
	c.mu.Unlock()

	return c
}

// Empty returns true if the composed error has no nested errors.
func (c *ConcurrentComposedError) Empty() bool {
	c.mu.Lock()
	e := c.ComposedError.Empty()
	c.mu.Unlock()

	return e
}

// Error implements error.Error
func (c *ConcurrentComposedError) Error() string {
	c.mu.Lock()
	e := c.ComposedError.Error()
	c.mu.Unlock()

	return e
}
