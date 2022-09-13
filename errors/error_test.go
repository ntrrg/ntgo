// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors_test

import (
	"errors"
	"testing"

	nterrors "go.ntrrg.dev/ntgo/errors"
)

var cases = []struct {
	label  string
	code   string
	reason string
	err    error
	want   string
}{
	{
		label:  "Simple",
		code:   "500",
		reason: "internal server error",
		want:   "[500] internal server error",
	},

	{
		label:  "Scoped",
		code:   "storage/tx/done",
		reason: "transaction has already been committed or rolled back",
		want:   "[storage/tx/done] transaction has already been committed or rolled back",
	},

	{
		label:  "Wrapped",
		code:   "net/http",
		reason: "can't start server",
		err:    errors.New("listen tcp :80: bind: address already in use"),
		want:   "[net/http] can't start server: listen tcp :80: bind: address already in use",
	},
}

func TestNew(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		c := c

		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			err := nterrors.New(c.code, c.reason)

			if code := err.Code(); code != c.code {
				t.Errorf("invalid code. got: %q, want: %q", code, c.code)
			}

			if reason := err.Reason(); reason != c.reason {
				t.Errorf("invalid reason. got: %q, want: %q", reason, c.reason)
			}
		})
	}
}

func TestError_Clone(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		c := c

		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			orig := nterrors.New(c.code, c.reason)
			cpy := orig.Clone()

			if orig == cpy {
				t.Error("cloned error is the same as source.")
			}

			if cpy.Code() != orig.Code() {
				msg := "invalid code cloned. got: %q, want: %q"
				t.Errorf(msg, cpy.Code(), orig.Code())
			}

			if cpy.Reason() != orig.Reason() {
				msg := "invalid reason cloned. got: %q, want: %q"
				t.Errorf(msg, cpy.Reason(), orig.Reason())
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		c := c

		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			err := nterrors.New(c.code, c.reason)
			if c.err != nil {
				err = err.Wrap(c.err)
			}

			if got := err.Error(); got != c.want {
				t.Errorf("invalid error. got: %q, want: %q", got, c.want)
			}
		})
	}
}
