// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"

	. "github.com/ntrrg/ntgo/math/arithmetic"
)

func ExampleAdd() {
	// See implementation_test.go
	var x, y Operand = "a", "b"

	r := Add(x, y)
	fmt.Println(r)
	// Output: 2
}
