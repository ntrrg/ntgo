// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	"nt.web.ve/go/ntgo/generics/arithmetic"
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

func ExampleOperander() {
	x := BytesSum("hello")
	fmt.Println(x.Val())

	fmt.Println(arithmetic.Add(x, BytesSum("world")))
	fmt.Println(arithmetic.Sub(x, 32))
	fmt.Println(arithmetic.Mul(x, "world"))
	fmt.Println(arithmetic.Div(x, []byte{'M', 'A'}))

	fmt.Println(arithmetic.Add(x))
	fmt.Println(arithmetic.Sub(x))
	fmt.Println(arithmetic.Mul(x))
	fmt.Println(arithmetic.Div(x))

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
