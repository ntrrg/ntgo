// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors_test

import (
	"errors"
	"testing"

	nterrors "go.ntrrg.dev/ntgo/errors"
)

func TestAll(t *testing.T) {
	t.Parallel()

	err := nterrors.New("test-all", "test All")

	if nterrors.All(err) {
		t.Error("equality with no targets.")
	}

	targets := []error{
		errors.New("this is an error for All"),
		errors.New("this is other error for All"),
		errors.New("this is another error for All"),
	}

	if nterrors.All(err, targets...) {
		t.Errorf("equality with non Error.\n\t%q == %q", err, targets)
	}

	targets = []error{err, nterrors.New("test-other-all", "test All"), err}
	if nterrors.All(err, targets...) {
		t.Errorf("equality with different error code.\n\t%q == %q", err, targets)
	}

	targets = []error{err, nterrors.New("test-other-all", "test other"), err}
	if nterrors.All(err, targets...) {
		t.Errorf("equality with different error.\n\t%q == %q", err, targets)
	}

	targets = []error{err, nterrors.New(err.Code(), "test other All"), err}
	if !nterrors.All(err, targets...) {
		t.Errorf("inequality with equal error code.\n\t%q != %q", err, targets)
	}

	targets = []error{err, err, err}
	if !nterrors.All(err, targets...) {
		t.Errorf("inequality with equal Error.\n\t%q != %q", err, targets)
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	err := nterrors.New("test-any", "test Any")

	targets := []error{
		errors.New("this is an error for Any"),
		errors.New("this is another error for Any"),
	}

	if nterrors.Any(err, targets...) {
		t.Errorf("equality with non Error.\n\t%q == %q", err, targets)
	}

	targets = append(targets, nterrors.New("test-other-any", "test Any"))
	if nterrors.Any(err, targets...) {
		t.Errorf("equality with different error code.\n\t%q == %q", err, targets)
	}

	targets = append(targets, nterrors.New("test-other-any", "test other"))
	if nterrors.Any(err, targets...) {
		t.Errorf("equality with different error.\n\t%q == %q", err, targets)
	}

	targets = append(targets, nterrors.New(err.Code(), "test other Any"))
	if !nterrors.Any(err, targets...) {
		t.Errorf("inequality with equal error code.\n\t%q != %q", err, targets)
	}

	targets = append(targets, nterrors.New(err.Code(), err.Reason()))
	if !nterrors.Any(err, targets...) {
		t.Errorf("inequality with equal Error.\n\t%q != %q", err, targets)
	}
}

func TestIs(t *testing.T) {
	t.Parallel()

	err := nterrors.New("test-error-is", "test Error.Is")
	other := errors.New("this is an error for Error.Is")

	if errors.Is(err, nil) {
		t.Error("equality with nil error.")
	}

	if errors.Is(err, &nterrors.Error{}) {
		t.Error("equality with empty error.")
	}

	if errors.Is(err, other) {
		t.Errorf("equality with non Error.\n\t%q == %q", err, other)
	}

	other = nterrors.New("test-other-error-is", "test Error.Is")
	if errors.Is(err, other) {
		t.Errorf("equality with different error code.\n\t%q == %q", err, other)
	}

	other = nterrors.New("test-other-error-is", "test other")
	if errors.Is(err, other) {
		t.Errorf("equality with different error.\n\t%q == %q", err, other)
	}

	other = nterrors.New(err.Code(), "test other Error.Is")
	if !errors.Is(err, other) {
		t.Errorf("inequality with equal error code.\n\t%q != %q", err, other)
	}

	other = nterrors.New(err.Code(), err.Reason())
	if !errors.Is(err, other) {
		t.Errorf("inequality with equal Error.\n\t%q != %q", err, other)
	}
}

func TestOf(t *testing.T) {
	t.Parallel()

	err := nterrors.New("test-error-of", "test Error.Of")
	other := nterrors.New("test-error-of_other", "test other Error.Of")
	nerr := nterrors.New(err.Code()+"/sub", "sub"+err.Reason())
	stderr := errors.New("[test-error-of] test Error.Of")

	if nterrors.Of(err, nil) {
		t.Error("success with nil error.")
	}

	if nterrors.Of(err, stderr) {
		t.Error("success with stdlib error as target.")
	}

	if nterrors.Of(stderr, err) {
		t.Error("success with stdlib error as source.")
	}

	if nterrors.Of(err, &nterrors.Error{}) {
		t.Error("success with empty error.")
	}

	if nterrors.Of(err, other) {
		t.Errorf("success with unrelated error.\n\t%q -> %q", other, err)
	}

	if nterrors.Of(err, nerr) {
		t.Errorf("success with child error.\n\t%q -> %q", nerr, err)
	}

	if !nterrors.Of(nerr, err) {
		t.Errorf("failure with parent error.\n\t%q -> %q", err, nerr)
	}
}

func TestWrapping(t *testing.T) {
	t.Parallel()

	var err error = nil

	if errs := nterrors.UnwrapAll(err); len(errs) != 0 {
		t.Errorf("unwrapped errors from nil. got: %q", errs)
	}

	err = errors.New("test wrapping")

	if errs := nterrors.UnwrapAll(err); len(errs) != 0 {
		t.Errorf("unwrapped errors from single stdlib error. got: %q", errs)
	}

	err = nterrors.New("test-wrapping", "test wrapping")

	if errs := nterrors.UnwrapAll(err); len(errs) != 0 {
		t.Errorf("unwrapped errors from single Error. got: %q", errs)
	}

	stderr := errors.New("testing wrapped stdlib")
	err = nterrors.Wrap(err, stderr)

	if errs := nterrors.UnwrapAll(err); len(errs) != 1 {
		t.Errorf("unwrapped more errors. got: %q, from: %q", errs, err)
	}

	err = nterrors.Wrap(stderr, err)

	if errs := nterrors.UnwrapAll(err); len(errs) != 2 {
		t.Errorf("unwrapped more errors. got: %q, from: %q", errs, err)
	}

	err = nterrors.Wrap(nil, err)

	if errs := nterrors.UnwrapAll(err); len(errs) != 2 {
		t.Errorf("unwrapped more errors from nil. got: %q, from: %q", errs, err)
	}

	err = nterrors.WrapAll()

	if errs := nterrors.UnwrapAll(err); len(errs) != 0 {
		t.Errorf("unwrapped errors from none. got: %q", errs)
	}

	err = nterrors.WrapAll(nil)

	if errs := nterrors.UnwrapAll(err); len(errs) != 0 {
		t.Errorf("unwrapped errors from nil. got: %q", errs)
	}

	err = nterrors.WrapAll(
		errors.New("parent"),
		nil,
		errors.New("first child"),
		errors.New("second child"),
		nil,
		nterrors.New("third-child", "third child"),
	)

	if errs := nterrors.UnwrapAll(err); len(errs) != 3 {
		t.Errorf("unwrapped nil errors. got: %q, from: %q", errs, err)
	}
}
