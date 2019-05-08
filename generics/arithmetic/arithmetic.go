// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic

import (
	"reflect"
)

// Operander is the interface that wraps the arithmetic representation method.
// It is useful for adding custom behavior to named types when GetVal processes
// it, other wise, the underlying type is obtained and follows the extraction
// rules.
//
// Val returns the variable's arithmetic representation (float64).
type Operander interface {
	Val() float64
}

// Add gets any number of elements and returns their addition.
func Add(operanders ...interface{}) float64 {
	result := GetVal(operanders[0])

	for _, v := range operanders[1:] {
		result += GetVal(v)
	}

	return result
}

// Div gets any number of elements and returns their division.
func Div(operanders ...interface{}) float64 {
	result := GetVal(operanders[0])

	for _, v := range operanders[1:] {
		result /= GetVal(v)
	}

	return result
}

// Eq gets any number of elements and checks if they are equals.
func Eq(operanders ...interface{}) bool {
	x := GetVal(operanders[0])

	for _, v := range operanders[1:] {
		if x != GetVal(v) {
			return false
		}
	}

	return true
}

// Mul gets any number of elements and returns their multiplication.
func Mul(operanders ...interface{}) float64 {
	result := GetVal(operanders[0])

	for _, v := range operanders[1:] {
		result *= GetVal(v)
	}

	return result
}

// Ne gets any number of elements and checks if they are differents.
func Ne(operanders ...interface{}) bool {
	s := make(map[float64]struct{})

	for _, v := range operanders {
		ar := GetVal(v)

		if _, ok := s[ar]; ok {
			return false
		}

		s[ar] = struct{}{}
	}

	return true
}

// Sub gets any number of elements and returns their subtraction.
func Sub(operanders ...interface{}) float64 {
	result := GetVal(operanders[0])

	for _, v := range operanders[1:] {
		result -= GetVal(v)
	}

	return result
}

/*
GetVal extracts the arithmetic representation from any type. It is ruled by the
value extraction rules.

Value extraction rules

1. Any element that satisfies the Operander interface will obtain its value
from the Val method.

2. Any element with a named type that doesn't satisfies the Operander interface
will obtain its value from its underlying type.

3. Boolean elements with a true value will be represented as 1, for false
values they will be 0.

4. Numeric elements (int, int8, int16, int32, int64, uint, uint8, uint16,
uint32, uint64, float32, float64, complex64, complex128, byte, rune) will be
converted to float64, but complex numbers will be represented by their real
part in float64 form.

5. Composed (arrays, maps, slices, strings, structs) and channel elements will
be represented by their length (or their number of fields for structs).

6. Any other element will be 0.
*/
func GetVal(operander interface{}) float64 {
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
