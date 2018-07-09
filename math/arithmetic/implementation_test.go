// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

// Operand is a simple string implementation of Operander.
type Operand string

func (o Operand) Val() float64 {
	return float64(len(o))
}
