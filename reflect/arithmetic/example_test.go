// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	a "go.ntrrg.dev/ntgo/reflect/arithmetic"
)

// BytesSum is a simple string implementation of arithmetic.Operander. Its
// arithmetic representation is the sum of all its bytes values.
type BytesSum string

func (o BytesSum) Val() (r float64) {
	for _, v := range o {
		r += float64(v)
	}

	return r
}

func Example() {
	x := BytesSum("hello")
	fmt.Println(x.Val())

	fmt.Println(a.Add(x, BytesSum("world")))
	fmt.Println(a.Sub(x, 32))
	fmt.Println(a.Mul(x, "world"))
	fmt.Println(a.Div(x, []byte{'M', 'A'}))

	fmt.Println(a.Add(x))
	fmt.Println(a.Sub(x))
	fmt.Println(a.Mul(x))
	fmt.Println(a.Div(x))

	// Output:
	// 532
	// 1084
	// 500
	// 2660
	// 266
	// 532
	// 532
	// 532
	// 532
}
