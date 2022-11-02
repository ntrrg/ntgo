// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors

import (
	"errors"
)

// Parsing errors.
var (
	ErrInvalidSyntax = New(Err.Code()+"/parse", "invalid error message")

	ErrEmptyMessage = New(
		ErrInvalidSyntax.Code()+"/empty",
		"empty error message",
	)

	// Code errors.

	ErrInvalidCode = New(ErrInvalidSyntax.Code()+"/code", "invalid code")

	ErrInvalidCodeChar = New(
		ErrInvalidCode.Code()+"/char",
		"invalid code character",
	)

	ErrNoCode = New(ErrInvalidCode.Code()+"/none", "error message has no code")

	// Reason errors.

	ErrInvalidReason = New(ErrInvalidSyntax.Code()+"/reason", "invalid reason")

	ErrNoReason = New(
		ErrInvalidReason.Code()+"/none",
		"error message has no reason",
	)

	ErrNoReasonSeparator = New(
		ErrInvalidReason.Code()+"/no-separator",
		"reason has no separator",
	)
)

// MustParse is like Parse, but panics if there is some syntax error. This is
// an utility function for package level errors initialization.
func MustParse(msg string) *Error {
	e, err := Parse(msg)
	if err != nil {
		panic(ErrInvalidSyntax.Wrap(err))
	}

	return e
}

// Parse recreates an Error from the given error message, if valid.
func Parse(msg string) (*Error, error) {
	if len(msg) == 0 {
		return nil, ErrEmptyMessage
	}

	code, msg, errPC := parseCode(msg)
	if errPC != nil {
		return nil, ErrInvalidCode.Wrap(errPC)
	}

	reason, msg, errPR := parseReason(msg)
	if errPR != nil {
		return nil, ErrInvalidReason.Wrap(errPR)
	}

	e := New(code, reason)

	if len(msg) == 0 {
		return e, nil
	}

	if werr, err := Parse(msg); err == nil {
		e.err = werr
	} else {
		e.err = errors.New(msg)
	}

	return e, nil
}

/**
 * Error
 */

// New creates an error based on e, appends code to e code and overrides its
// reason. This is an utility method for package level errors initialization,
// thus, providing bad syntax panics.
func (e *Error) New(code, reason string) *Error {
	err := New(e.Code()+"/"+code, reason)
	return MustParse(err.Error())
}

/**
 * Helpers
 */

func isValidCodeChar(r byte) bool {
	switch {
	case r >= 'a' && r <= 'z':
		return true
	case r >= '0' && r <= '9':
		return true
	case r == '_' || r == '-' || r == '.':
		return true
	default:
		return false
	}
}

func parseCode(msg string) (code, nmsg string, err error) {
	if len(msg) < len("[x]") {
		return "", msg, ErrNoCode
	}

	if msg[0] != '[' || msg[1] == ']' {
		return "", msg, ErrNoCode
	}

	nmsg = msg[1:]

	for i, l := 0, len(nmsg); i < l; i++ {
		b := nmsg[i]

		switch {
		case b == ']':
			return nmsg[:i], nmsg[i+1:], nil

		case b != '/' && !isValidCodeChar(b):
			err := errors.New("invalid byte '" + string(b) + "'")
			return "", msg, ErrInvalidCodeChar.Wrap(err)
		}
	}

	return "", msg, ErrNoCode
}

func parseReason(msg string) (reason, nmsg string, err error) {
	if len(msg) < len(" x") {
		return "", msg, ErrNoReason
	}

	if msg[0] != ' ' {
		return "", msg, ErrNoReasonSeparator
	}

	nmsg = msg[1:]

	for i, l := 1, len(nmsg); i < l; i++ {
		if nmsg[i] != ':' {
			continue
		}

		if nmsg[i-1] == ' ' {
			continue
		}

		if i+1 < len(nmsg) && nmsg[i+1] == ' ' {
			return nmsg[:i], nmsg[i+2:], nil
		}
	}

	return nmsg, "", nil
}
