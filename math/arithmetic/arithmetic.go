// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

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

// Add gets any number of Operanders and returns their addition.
func Add(o ...Operander) float64 {
	result := o[0].Val()

	for _, v := range o[1:] {
		result += v.Val()
	}

	return result
}

// Div gets any number of Operanders and returns their division.
func Div(o ...Operander) float64 {
	result := o[0].Val()

	for _, v := range o[1:] {
		result /= v.Val()
	}

	return result
}

// Mul gets any number of Operanders and returns their multiplication.
func Mul(o ...Operander) float64 {
	result := o[0].Val()

	for _, v := range o[1:] {
		result *= v.Val()
	}

	return result
}

// Sub gets any number of Operanders and returns their subtraction.
func Sub(o ...Operander) float64 {
	result := o[0].Val()

	for _, v := range o[1:] {
		result -= v.Val()
	}

	return result
}
