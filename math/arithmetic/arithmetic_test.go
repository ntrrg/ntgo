// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"testing"

	. "github.com/ntrrg/ntgo/math/arithmetic"
)

// See implementation_test.go

func TestAdd(t *testing.T) {
	cases := []struct {
		x, y Operand
		want float64
	}{
		{"a", "b", 2},
		{"abcd", "xyz", 7},
		{"m", " a", 3},
		{"", "", 0},
	}

	for _, c := range cases {
		r := Add(c.x, c.y)

		if r != c.want {
			t.Errorf("%v.Add(%v) == %v, want %v", c.x, c.y, r, c.want)
		}
	}
}
