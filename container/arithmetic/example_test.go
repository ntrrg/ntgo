// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	"github.com/ntrrg/ntgo/container/arithmetic"
)

func Example() {
	x := arithmetic.Add("Miguel", "Angel", arithmetic.Sub(5, []int{1, 2, 3}))
	y := arithmetic.Mul(2, "four") + arithmetic.Div(6, "two")
	r := x - y
	fmt.Println(r)
	// Output: 3
}

func ExampleAdd() {
	r := arithmetic.Add(true, false)
	fmt.Println(r)
	// Output: 1
}

func ExampleDiv() {
	r := arithmetic.Div(12, 6.0)
	fmt.Println(r)
	// Output: 2
}

func ExampleMul() {
	r := arithmetic.Mul(1+2i, func() {})
	fmt.Println(r)
	// Output: 0
}

func ExampleSub() {
	r := arithmetic.Sub([3]string{"hello", ", ", "world!"}, []int{1, 2})
	fmt.Println(r)
	// Output: 1
}
