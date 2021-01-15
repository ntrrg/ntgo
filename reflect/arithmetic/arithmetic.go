// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic

import (
	"reflect"
)

// An Operander is a value that can be represented as an arithmetic operand.
type Operander interface {
	Val() float64
}

// Add gets any number of elements and returns their addition.
func Add(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := Val(operanders[0])

	for _, v := range operanders[1:] {
		result += Val(v)
	}

	return result
}

// Div gets any number of elements and returns their division.
func Div(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := Val(operanders[0])

	for _, v := range operanders[1:] {
		result /= Val(v)
	}

	return result
}

// Eq gets any number of elements and checks if they are equals.
func Eq(operanders ...interface{}) bool {
	if len(operanders) < 2 {
		return true
	}

	x := Val(operanders[0])

	for _, v := range operanders[1:] {
		if x != Val(v) {
			return false
		}
	}

	return true
}

// Mul gets any number of elements and returns their multiplication.
func Mul(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := Val(operanders[0])

	for _, v := range operanders[1:] {
		result *= Val(v)
	}

	return result
}

// Ne gets any number of elements and checks if they are differents.
func Ne(operanders ...interface{}) bool {
	if len(operanders) < 2 {
		return false
	}

	s := make(map[float64]struct{}, len(operanders))

	for _, v := range operanders {
		val := Val(v)

		if _, ok := s[val]; ok {
			return false
		}

		s[val] = struct{}{}
	}

	return true
}

// Sub gets any number of elements and returns their subtraction.
func Sub(operanders ...interface{}) float64 {
	if len(operanders) < 1 {
		return 0
	}

	result := Val(operanders[0])

	for _, v := range operanders[1:] {
		result -= Val(v)
	}

	return result
}

// Val extracts the arithmetic representation from any type. It is ruled by the
// value extraction rules.
func Val(operander interface{}) float64 {
	if x, ok := operander.(Operander); ok {
		return x.Val()
	}

	x := reflect.ValueOf(operander)

	// nolint:exhaustive
	switch x.Kind() {
	case reflect.Bool:
		if x.Bool() {
			return 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return float64(x.Uint())
	case reflect.Float32, reflect.Float64:
		return x.Float()
	case reflect.Complex64, reflect.Complex128:
		y := x.Complex()
		return real(y)
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return float64(x.Len())
	case reflect.Struct:
		return float64(x.NumField())
	}

	return 0
}
