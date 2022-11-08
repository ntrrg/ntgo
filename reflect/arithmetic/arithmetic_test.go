// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"testing"

	"go.ntrrg.dev/ntgo/reflect/arithmetic"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	if arithmetic.Add() != 0 {
		t.Errorf("Add with zero operanders should be 0")
	}

	if arithmetic.Add(1) != 1 {
		t.Errorf("Add with one operander should be the same operander")
	}
}

func TestDiv(t *testing.T) {
	t.Parallel()

	if arithmetic.Div() != 0 {
		t.Errorf("Div with zero operanders should be 0")
	}

	if arithmetic.Div(1) != 1 {
		t.Errorf("Div with one operander should be the same operander")
	}
}

func TestEq(t *testing.T) {
	t.Parallel()

	if !arithmetic.Eq() || !arithmetic.Eq('M') ||
		!arithmetic.Eq(true, 1, []byte{'M'}) ||
		arithmetic.Eq(false, true, []byte{'M', 'A'}) ||
		!arithmetic.Eq(false, 0, []byte{}, struct{}{}, func() {}) {
		t.Errorf("equality is not working as it should")
	}
}

func TestMul(t *testing.T) {
	t.Parallel()

	if arithmetic.Mul() != 0 {
		t.Errorf("Mul with zero operanders should be 0")
	}

	if arithmetic.Mul(1) != 1 {
		t.Errorf("Mul with one operander should be the same operander")
	}
}

func TestNe(t *testing.T) {
	t.Parallel()

	if arithmetic.Ne() || arithmetic.Ne('M') ||
		arithmetic.Ne(true, 1, []byte{'M'}) ||
		!arithmetic.Ne(false, true, []byte{'M', 'A'}) ||
		arithmetic.Ne(false, 0, []byte{}, struct{}{}, func() {}) {
		t.Errorf("inequality is not working as it should")
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	if arithmetic.Sub() != 0 {
		t.Errorf("Sub with zero operanders should be 0")
	}

	if arithmetic.Sub(1) != 1 {
		t.Errorf("Sub with one operander should be the same operander")
	}
}

func TestVal(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   any
		want float64
	}{
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
		got := arithmetic.Val(c.in)

		if got != c.want {
			t.Errorf("%v == %v, want %v", c.in, got, c.want)
		}
	}
}
