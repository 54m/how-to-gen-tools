// Code generated by volcago. DO NOT EDIT.
// generated version: v1.0.0
package repository

import (
	"fmt"
	"strings"
)

// MultiError - multi error interface
type MultiError interface {
	GetIndex() int
	GetError() error
	Error() string
}

// NewMultiError - constructor
func NewMultiError(index int, err error) MultiError {
	return &multiError{
		index: index,
		err:   err,
	}
}

// compile time interface checks
var _ MultiError = (*multiError)(nil)
var _ error = (*multiError)(nil)

// multiError - multi error
type multiError struct {
	index int
	err   error
}

// GetIndex - getter index
func (e *multiError) GetIndex() int {
	return e.index
}

// GetError - getter error
func (e *multiError) GetError() error {
	return e.err
}

// Error - multi error string
func (e *multiError) Error() string {
	return fmt.Sprintf("[index=%d, err=%v]", e.index, e.err)
}

// MultiErrors - multi errors
type MultiErrors []MultiError

// NewMultiErrors - constructor
func NewMultiErrors() MultiErrors {
	return make([]MultiError, 0)
}

// compile time interface checks
var _ error = (*MultiErrors)(nil)

// Error - multi errors string
func (errs MultiErrors) Error() string {
	msgs := make([]string, 0)
	for _, e := range errs {
		msgs = append(msgs, e.Error())
	}

	return fmt.Sprintf("multi errors: %s", strings.Join(msgs, ", "))
}
