// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"testing"

	"github.com/ntrrg/ntgo/math/arithmetic"
)

// For BytesSum see operander_example_test.go

func TestGetVal(t *testing.T) {
	cases := []struct {
		in   interface{}
		want float64
	}{
		{BytesSum("hello"), 532},
		{true, 1},
		{int(1), 1},
		{int8(2), 2},
		{int16(3), 3},
		{int32(4), 4},
		{int64(5), 5},
		{uint(6), 6},
		{uint8(7), 7},
		{uint16(8), 8},
		{uint32(9), 9},
		{uint64(10), 10},
		{float32(11.12e3), 11.12e3},
		{float64(14.15e-6), 14.15e-6},
		{complex64(17 + 18i), 17},
		{complex128(19 - 20i), 19},
		{[3]int{21, 22, 23}, 3},
		{make(chan<- string), 0},
		{map[string]int{"24": 25, "26": 27, "28": 29}, 3},
		{[]int{30, 31, 32, 33, 34}, 5},
		{"hello, world!", 13},
		{struct{ name string }{"Miguel"}, 1},
		{'M', 77},
		{'😄', 128516},
		{false, 0},
		{[]string{}, 0},
		{make(map[bool]string), 0},
		{"", 0},
		{struct{}{}, 0},
		{func() {}, 0},
	}

	for _, c := range cases {
		got := arithmetic.GetVal(c.in)

		if got != c.want {
			t.Errorf("%v == %v, want %v", c.in, got, c.want)
		}
	}
}
