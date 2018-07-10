// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	a "github.com/ntrrg/ntgo/math/arithmetic"
)

// See common_test.go

func ExampleAdd() {
	var x, y Operand = "a", "b"

	r := a.Add(x, y)
	fmt.Println(r)
	// Output: 2
}

func ExampleDiv() {
	var x, y Operand = "abcdef", "xy"

	r := a.Div(x, y)
	fmt.Println(r)
	// Output: 3
}

func ExampleMul() {
	var x, y Operand = "abc", "xyz"

	r := a.Mul(x, y)
	fmt.Println(r)
	// Output: 9
}

func ExampleSub() {
	var x, y Operand = "aeiou", "bcdfg"

	r := a.Sub(x, y)
	fmt.Println(r)
	// Output: 0
}
