// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors

// Error records an error during any operation.
type Error struct {
	code   string
	reason string
	err    error
}

// New creates an Error with the given data.
func New(code, reason string) *Error {
	return &Error{code: code, reason: reason}
}

// Clone returns a copy of e.
func (e *Error) Clone() *Error {
	return &Error{code: e.code, reason: e.reason}
}

// Code retruns e unique identifier.
func (e *Error) Code() string {
	return e.code
}

// Error implements the error interface.
func (e *Error) Error() string {
	err := "[" + e.code + "] " + e.reason

	if e.err != nil {
		err += ": " + e.err.Error()
	}

	return err
}

// Is customizes the functionality of errors.Is. Every error is identified by
// its code, which means, Is will return true even if their reason is
// different, but their code is the same.
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error) //nolint:errorlint
	if !ok {
		return false
	}

	return e.code == t.code
}

// Of reports if e was created from target.
func (e *Error) Of(target *Error) bool {
	return hasPrefix(e.code, target.code+"/")
}

// Reason retruns e friendly description.
func (e *Error) Reason() string {
	return e.reason
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	return e.err
}

// Wrap returns a copy of e that wraps the given error.
func (e *Error) Wrap(err error) *Error {
	ne := e.Clone()
	ne.err = err

	return ne
}

/**
 * Helpers
 */

func hasPrefix(s, prefix string) bool {
	l := len(prefix)

	if l == 0 || len(s) < l {
		return false
	}

	return s[:l] == prefix
}
