// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"testing"

	"nt.web.ve/go/ntgo/reflect/arithmetic"
)

func BenchmarkGetVal(b *testing.B) {
	cases := []struct {
		name string
		val  interface{}
	}{
		{"Bool", true},
		{"Int", int(1)},
		{"Int8", int8(2)},
		{"Int16", int16(3)},
		{"Int32", int32(4)},
		{"Int64", int64(5)},
		{"Uint", uint(6)},
		{"Uint8", uint8(7)},
		{"Uint16", uint16(8)},
		{"Uint32", uint32(9)},
		{"Uint64", uint64(10)},
		{"Float32", float32(11.12e3)},
		{"Float64", float64(14.15e-6)},
		{"Complex64", complex64(17 + 18i)},
		{"Complex128", complex128(19 - 20i)},
		{"Array", [3]int{21, 22, 23}},
		{"Channel", make(chan<- string)},
		{"Map", map[string]int{"24": 25, "26": 27, "28": 29}},
		{"Slice", []int{30, 31, 32, 33, 34}},
		{"String", "hello, world!"},
		{"Function", func() {}},
	}

	for _, c := range cases {
		c := c

		b.Run(c.name, func(b *testing.B) {
			for i := 0; i <= b.N; i++ {
				arithmetic.GetVal(c.val)
			}
		})
	}
}
