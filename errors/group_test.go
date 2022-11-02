// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors_test

import (
	"errors"
	"testing"

	nterrors "go.ntrrg.dev/ntgo/errors"
)

func TestGrouping(t *testing.T) {
	t.Parallel()

	msg1 := "this is an error for Group"
	msg2 := "this is other error for Group"
	msg3 := "this is another error for Group"

	errs := []error{
		errors.New(msg1),
		errors.New(msg2),
		errors.New(msg3),
	}

	g := nterrors.Group(errs...)

	want := "* " + msg1 + "; * " + msg2 + "; * " + msg3
	got := g.Error()

	if got != want {
		t.Errorf("invalid error group. got: %q, want: %q", got, want)
	}

	for _, err := range errs {
		if !errors.Is(g, err) {
			t.Errorf("error not in group. err: %q, group: %q", err, g)
		}
	}

	targets := nterrors.Split(g)

	for i := range errs {
		want := errs[i].Error()
		got := targets[i].Error()

		if want != got {
			t.Errorf("invalid error in group. got: %q, want: %q", got, want)
		}
	}

	err := errors.New("this is a single error for Split")
	errs2 := nterrors.Split(err)

	if l := len(errs2); l != 1 {
		t.Fatalf("single error splited incorrectly. got: %d elements", l)
	}

	if got := errs2[0]; !errors.Is(got, err) {
		t.Errorf("invalid error in splited group. got: %q, want: %q", got, err)
	}
}
