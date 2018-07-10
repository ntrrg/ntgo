// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	a "github.com/ntrrg/ntgo/math/arithmetic"
)

// Operand is a simple string implementation of Operander and SelfOperander.
type Operand string

func (o Operand) Val() float64 {
	return float64(len(o))
}

func Operanders(o []Operand) []a.Operander {
	no := make([]a.Operander, len(o))

	for i, v := range o {
		no[i] = v
	}

	return no
}
