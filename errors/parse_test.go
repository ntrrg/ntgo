// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package errors_test

import (
	"errors"
	"strings"
	"testing"

	nterrors "go.ntrrg.dev/ntgo/errors"
)

func TestError_New(t *testing.T) {
	t.Parallel()

	code := "test-error-new"
	reason := "test Error.New"

	err := nterrors.New(code, reason)
	nerr := err.New("new", "new-"+reason)

	if ncode := nerr.Code(); ncode != code+"/new" {
		t.Errorf("invalid code. got: %q, want: %q", ncode, code+"/new")
	}

	if nreason := nerr.Reason(); nreason != "new-"+reason {
		t.Errorf("invalid reason. got: %q, want: %q", nreason, "new-"+reason)
	}
}

var parseCases = []struct {
	label string
	msg   string
	want  error
}{
	{
		label: "Simple",
		msg:   "[500] internal server error",
		want:  nterrors.New("500", "internal server error"),
	},

	{
		label: "Scoped",
		msg:   "[storage/tx/done] transaction has already been committed or rolled back",

		want: nterrors.New(
			"storage/tx/done",
			"transaction has already been committed or rolled back",
		),
	},

	{
		label: "Wrapped",
		msg:   "[net/http] can't start server: listen tcp :80: bind: address already in use",

		want: nterrors.New(
			"net/http", "can't start server",
		).Wrap(errors.New("listen tcp :80: bind: address already in use")),
	},

	{
		label: "WrappedPartialSeparator",
		msg:   "[wrapped-partial-separator] test wrapped partial separator:",

		want: nterrors.New(
			"wrapped-partial-separator", "test wrapped partial separator:",
		),
	},

	{
		label: "WrappedSeparatorBegining",
		msg:   "[wrapped-separator_begining] : test wrapped separator",

		want: nterrors.New(
			"wrapped-separator_begining", ": test wrapped separator",
		),
	},

	{
		label: "WrappedSeparatorInside",
		msg:   "[wrapped-separator_inside] test wrapped : separator",

		want: nterrors.New(
			"wrapped-separator_inside", "test wrapped : separator",
		),
	},

	{
		label: "WrappedSeparatorEnd",
		msg:   "[wrapped-separator_end] test wrapped separator:",
		want:  nterrors.New("wrapped-separator_end", "test wrapped separator:"),
	},

	{
		label: "WrappedEmpty",
		msg:   "[wrapped-empty] test wrapped empty: ",
		want:  nterrors.New("wrapped-empty", "test wrapped empty"),
	},

	{
		label: "DeepWrapped",
		msg:   "[test-wrap/deep/top] top level: [test-wrap/deep/mid] mid level: low level",

		want: nterrors.New("test-wrap/deep/top", "top level").Wrap(
			nterrors.New("test-wrap/deep/mid", "mid level").Wrap(
				errors.New("low level"),
			),
		),
	},
}

var parseErrorCases = []struct {
	label string
	msg   string
	want  []error
}{
	{
		label: "Empty",
		msg:   "",
		want:  []error{nterrors.ErrEmptyMessage},
	},

	{
		label: "NoCode",
		msg:   "test no code",
		want:  []error{nterrors.ErrInvalidCode, nterrors.ErrNoCode},
	},

	{
		label: "ShortCode",
		msg:   "[]",
		want:  []error{nterrors.ErrInvalidCode, nterrors.ErrNoCode},
	},

	{
		label: "NoClosedCode",
		msg:   "[test-no-closed-code",
		want:  []error{nterrors.ErrInvalidCode, nterrors.ErrNoCode},
	},

	{
		label: "BadCode",
		msg:   "[bad+code] test bad code",
		want:  []error{nterrors.ErrInvalidCode, nterrors.ErrInvalidCodeChar},
	},

	{
		label: "NoReason",
		msg:   "[test-no-reason]",
		want:  []error{nterrors.ErrInvalidReason, nterrors.ErrNoReason},
	},

	{
		label: "EmptyReason",
		msg:   "[test-empty-reason] ",
		want:  []error{nterrors.ErrInvalidReason, nterrors.ErrNoReason},
	},

	{
		label: "BadReason",
		msg:   "[bad-reason]test bad reason",
		want:  []error{nterrors.ErrInvalidReason, nterrors.ErrNoReasonSeparator},
	},
}

func TestMustParse(t *testing.T) {
	t.Parallel()

	for _, c := range parseCases {
		c := c

		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			defer func() {
				err := recover()
				if err == nil {
					return
				}

				t.Error(err)
			}()

			got := nterrors.MustParse(c.msg)
			if got.Error() != c.want.Error() {
				t.Fatalf("invalid error. got: %q, want: %q", got, c.want)
			}

			ww := nterrors.UnwrapAll(c.want)
			gw := nterrors.UnwrapAll(got)

			if len(ww) != len(gw) {
				t.Errorf("invalid wrapped error. got: %q, want: %q", gw, ww)
			}
		})
	}

	t.Run("Errors", func(t *testing.T) {
		t.Parallel()

		for _, c := range parseErrorCases {
			c := c

			t.Run(c.label, func(t *testing.T) {
				t.Parallel()

				defer func() {
					erri := recover()
					if erri == nil {
						return
					}

					got, _ := erri.(error) // nolint:errcheck
					if !nterrors.All(got, append(c.want, nterrors.ErrInvalidSyntax)...) {
						t.Errorf("invalid error. got: %q, want: %q", got, c.want)
					}
				}()

				nterrors.MustParse(c.msg)
				t.Error("succeed")
			})
		}
	})
}

func TestParse(t *testing.T) {
	t.Parallel()

	for _, c := range parseCases {
		c := c

		t.Run(c.label, func(t *testing.T) {
			t.Parallel()

			var got, err error = nterrors.Parse(c.msg)
			if err != nil {
				got = err
			}

			if got.Error() != c.want.Error() {
				t.Errorf("invalid error. got: %q, want: %q", got, c.want)
			}

			ww := nterrors.UnwrapAll(c.want)
			gw := nterrors.UnwrapAll(got)

			if len(ww) != len(gw) {
				t.Errorf("invalid wrapped error. got: %q, want: %q", gw, ww)
			}
		})
	}

	t.Run("Errors", func(t *testing.T) {
		t.Parallel()

		for _, c := range parseErrorCases {
			c := c

			t.Run(c.label, func(t *testing.T) {
				t.Parallel()

				_, got := nterrors.Parse(c.msg)
				if got == nil {
					t.Fatal("succeed")
				}

				if !nterrors.All(got, c.want...) {
					t.Errorf("invalid error. got: %q, want: %q", got, c.want)
				}
			})
		}
	})
}

func FuzzParse(f *testing.F) {
	for _, c := range parseCases {
		f.Add(c.msg)
	}

	for _, c := range parseErrorCases {
		f.Add(c.msg)
	}

	f.Fuzz(func(t *testing.T, msg string) {
		for strings.HasSuffix(msg, ": ") {
			msg = strings.TrimSuffix(msg, ": ")
		}

		got, err := nterrors.Parse(msg)
		if err != nil {
			return
		}

		if got.Error() != msg {
			t.Errorf("invalid error. got: %q, want: %q", got, msg)
		}
	})
}

func BenchmarkParse(b *testing.B) {
	for _, c := range parseCases {
		b.Run("Stdlib/"+c.label, func(b *testing.B) {
			var err error

			for i := 0; i < b.N; i++ {
				err = errors.New(c.msg)
				if err == nil {
					b.Fatalf("no error created (%s)", c.msg)
				}
			}
		})

		b.Run("New/"+c.label, func(b *testing.B) {
			e, errP := nterrors.Parse(c.msg)
			if errP != nil {
				b.Fatal(errP)
			}

			code := e.Code()
			reason := e.Reason()

			b.ResetTimer()

			var err error

			for i := 0; i < b.N; i++ {
				err = nterrors.New(code, reason)
				_ = err
			}
		})

		b.Run("Parse/"+c.label, func(b *testing.B) {
			var err error

			for i := 0; i < b.N; i++ {
				_, err = nterrors.Parse(c.msg)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
