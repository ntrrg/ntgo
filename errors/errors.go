// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors

import (
	"errors"
	"fmt"
)

// Main error group.
var Err = New("ntgo/errors", "")

// All reports if err matches all errors in targets.
func All(err error, targets ...error) bool {
	if len(targets) == 0 {
		return false
	}

	for _, t := range targets {
		if !errors.Is(err, t) {
			return false
		}
	}

	return true
}

// Any reports if err matches any error from targets.
func Any(err error, targets ...error) bool {
	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}

	return false
}

// Of reports if err was created from target.
func Of(err, target error) bool {
	e, ok := err.(*Error) // nolint:errorlint
	if !ok {
		return false
	}

	t, ok := target.(*Error) // nolint:errorlint
	if !ok {
		return false
	}

	return e.Of(t)
}

// UnwrapAll returns all wrapped errors by err.
func UnwrapAll(err error) []error {
	var errs []error = nil

	for e := errors.Unwrap(err); e != nil; e = errors.Unwrap(e) {
		errs = append(errs, e)
	}

	return errs
}

// Wrap wraps target with err. If err is nil, target will be returned.
func Wrap(err, target error) error {
	if err == nil {
		return target
	}

	if target == nil {
		return err
	}

	switch e := err.(type) {
	case *Error:
		return e.Wrap(target)
	case interface{ Wrap(error) error }:
		return e.Wrap(target)
	}

	return fmt.Errorf("%v: %w", err, target)
}

// WrapAll wraps all given errors left to right.
func WrapAll(errs ...error) error {
	var err error = nil

	for i := len(errs) - 1; i >= 0; i-- {
		err = Wrap(errs[i], err)
	}

	return err
}
