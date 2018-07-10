// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"testing"

	a "github.com/ntrrg/ntgo/math/arithmetic"
)

// See common_test.go

type testCase struct {
	in   []Operand
	want float64
}

type testCases []testCase

func TestAdd(t *testing.T) {
	cases := testCases{
		{[]Operand{"a", "b"}, 2},
		{[]Operand{"abc", "lmn", "xyz"}, 9},
		{[]Operand{"hello", ", ", "world", "!"}, 13},
		{[]Operand{"ma", ""}, 2},
		{[]Operand{"", "rn"}, 2},
		{[]Operand{"", ""}, 0},
	}

	for _, c := range cases {
		got := a.Add(Operanders(c.in)...)

		if got != c.want {
			t.Errorf("Add(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestDiv(t *testing.T) {
	cases := testCases{
		{[]Operand{"a", "b"}, 1},
		{[]Operand{"abc", "lmn", "xyz"}, 0.3333333333333333},
		{[]Operand{"hello", ", ", "world", "!"}, 0.5},
		{[]Operand{"", "rn"}, 0},
	}

	for _, c := range cases {
		got := a.Div(Operanders(c.in)...)

		if got != c.want {
			t.Errorf("Div(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMul(t *testing.T) {
	cases := testCases{
		{[]Operand{"a", "b"}, 1},
		{[]Operand{"abc", "lmn", "xyz"}, 27},
		{[]Operand{"hello", ", ", "world", "!"}, 50},
		{[]Operand{"ma", ""}, 0},
		{[]Operand{"", "rn"}, 0},
		{[]Operand{"", ""}, 0},
	}

	for _, c := range cases {
		got := a.Mul(Operanders(c.in)...)

		if got != c.want {
			t.Errorf("Mul(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSub(t *testing.T) {
	cases := testCases{
		{[]Operand{"a", "b"}, 0},
		{[]Operand{"abc", "lmn", "xyz"}, -3},
		{[]Operand{"hello", ", ", "world", "!"}, -3},
		{[]Operand{"ma", ""}, 2},
		{[]Operand{"", "rn"}, -2},
		{[]Operand{"", ""}, 0},
	}

	for _, c := range cases {
		got := a.Sub(Operanders(c.in)...)

		if got != c.want {
			t.Errorf("Sub(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
