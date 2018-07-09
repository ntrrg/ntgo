// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

// Package arithmetic provides arithmetics operations for any type.
package arithmetic

// Identity constants
const (
	AdditiveIdentity       = 0
	MultiplicativeIdentity = 1
)

// Operander is the interface that wraps the arithmetic representation methods.
//
// Val returns the variable's arithmetic representation (float64).
type Operander interface {
	Val() float64
}

// Add gets any number of Operander and returns their addition.
func Add(operands ...Operander) float64 {
	result := float64(0)

	for _, v := range operands {
		if v.Val() == AdditiveIdentity {
			continue
		}

		result += v.Val()
	}

	return result
}
