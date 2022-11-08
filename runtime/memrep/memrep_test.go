// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

//go:build !race
// +build !race

package memrep_test

import (
	"bytes"
	"testing"

	ntruntime "go.ntrrg.dev/ntgo/runtime"
	"go.ntrrg.dev/ntgo/runtime/memrep"
)

func TestRead(t *testing.T) {
	t.Parallel()

	// Since some memory representations includes pointers, it is not possible to
	// predict its values, therefor setting want.Bytes to nil means that this
	// property will be ignored while comparing.
	cases := []struct {
		in   any
		want memrep.Representation
	}{
		{
			in: true,

			want: memrep.Representation{
				Type:  "bool",
				Size:  1,
				Bytes: []byte{1},
			},
		},

		{
			in: int32(32),

			want: memrep.Representation{
				Type:  "int32",
				Size:  4,
				Bytes: []byte{32, 0, 0, 0},
			},
		},

		{
			in: int64(0),

			want: memrep.Representation{
				Type:  "int64",
				Size:  8,
				Bytes: []byte{7: 0},
			},
		},

		{
			in: [5]byte{'h', 'e', 'l', 'l', 'o'},

			want: memrep.Representation{
				Type:  "[5]uint8",
				Size:  5,
				Bytes: []byte{'h', 'e', 'l', 'l', 'o'},
			},
		},

		{
			in: "hello, world",

			want: memrep.Representation{
				Type:  "string",
				Size:  16,
				Bytes: nil,
			},
		},

		{
			in: []byte("hello, world"),

			want: memrep.Representation{
				Type:  "[]uint8",
				Size:  24,
				Bytes: nil,
			},
		},

		{
			in: customStruct{a: true, b: 8, c: 16, d: 32, e: 64},

			want: memrep.Representation{
				Type:  "memrep_test.customStruct",
				Size:  16,
				Bytes: []byte{1, 8, 16, 0, 32, 8: 64, 15: 0},
			},
		},

		{
			in: &customStruct{},

			want: memrep.Representation{
				Type:  "*memrep_test.customStruct",
				Size:  ntruntime.WordSize(),
				Bytes: nil,
			},
		},

		{
			in: memrep.Representation{
				Type:  "bool",
				Size:  1,
				Bytes: []byte{1},
			},

			want: memrep.Representation{
				Type:  "memrep.Representation",
				Size:  48,
				Bytes: nil,
			},
		},
	}

	for _, c := range cases {
		var got memrep.Representation

		switch v := c.in.(type) {
		case []byte:
			got = memrep.Read(v)
		case [5]byte:
			got = memrep.Read(v)
		case bool:
			got = memrep.Read(v)
		case customStruct:
			got = memrep.Read(v)
		case *customStruct:
			got = memrep.Read(v)
		case int32:
			got = memrep.Read(v)
		case int64:
			got = memrep.Read(v)
		case memrep.Representation:
			got = memrep.Read(v)
		case string:
			got = memrep.Read(v)
		default:
			t.Fatalf("invalid type for test input. got: %+v (%[1]T)", v)
		}

		if !isValid(c.want, got) {
			t.Errorf("invalid representation.\ngot: %+v\nwant: %+v", got, c.want)
		}
	}
}

type customStruct struct {
	a bool
	b byte
	c int16
	d int32
	e int64
}

func TestReadArray(t *testing.T) {
	t.Parallel()

	in := []rune("hello, world!")

	want := memrep.Representation{
		Type:  "array from []int32",
		Size:  52,
		Bytes: nil, // Ignore when comparing.
	}

	if got := memrep.ReadArray(in); !isValid(want, got) {
		t.Errorf("invalid representation.\ngot: %+v\nwant: %+v", got, want)
	}
}

func isValid(a, b memrep.Representation) bool {
	if a.Type != b.Type {
		return false
	}

	if a.Size != b.Size {
		return false
	}

	if a.Bytes == nil || b.Bytes == nil {
		return true
	}

	if !bytes.Equal(a.Bytes, b.Bytes) {
		return false
	}

	return true
}
