// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic

import "reflect"

// Operander is the interface that wraps the arithmetic representation method.
//
// Val returns the variable's arithmetic representation (float64).
type Operander interface {
	Val() float64
}

// Add gets any number of elements and returns their addition.
func Add(operanders ...interface{}) float64 {
	result := getValue(operanders[0])

	for _, v := range operanders[1:] {
		result += getValue(v)
	}

	return result
}

// Div gets any number of elements and returns their division.
func Div(operanders ...interface{}) float64 {
	result := getValue(operanders[0])

	for _, v := range operanders[1:] {
		result /= getValue(v)
	}

	return result
}

// Mul gets any number of elements and returns their multiplication.
func Mul(operanders ...interface{}) float64 {
	result := getValue(operanders[0])

	for _, v := range operanders[1:] {
		result *= getValue(v)
	}

	return result
}

// Sub gets any number of elements and returns their subtraction.
func Sub(operanders ...interface{}) float64 {
	result := getValue(operanders[0])

	for _, v := range operanders[1:] {
		result -= getValue(v)
	}

	return result
}

func getValue(operander interface{}) float64 {
	if x, ok := operander.(Operander); ok {
		return x.Val()
	}

	x := reflect.ValueOf(operander)

	switch x.Kind() {
	case reflect.Bool:
		if x.Bool() {
			return 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(x.Uint())
	case reflect.Float32, reflect.Float64:
		return x.Float()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return float64(x.Len())
	}

	return 0
}
