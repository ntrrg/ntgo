// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	"github.com/ntrrg/ntgo/math/arithmetic"
)

// BytesSum is a simple string implementation of Operander. Its arithmetic
// representation is the sum of all its bytes.
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
	fmt.Println(arithmetic.Add(x, 8))
	// Output:
	// 532
	// 540
}
