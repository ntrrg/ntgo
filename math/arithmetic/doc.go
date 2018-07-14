// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

/*
Package arithmetic provides arithmetic operations for any type.

The four elemental operations (addition, subtraction, multiplication and
division) are defined as functions with their short names (Add, Sub, Mul and
Div). Since there are some data types without an arithmetical representation,
some rules are applied during the value extraction.

Value extraction rules

1. Any element that satisfies the Operander interface will obtain its value
from the Val method.

2. Any element with a named type that doesn't satisfies the Operander interface
will obtain its value from its underlying type.

3. Boolean elements with a true value will be represented as 1, for false
values they will be 0.

4. Numeric elements (int, int8, int16, int32, int64, uint, uint8, uint16,
uint32, uint64, float32, float64, complex64, complex128, byte, rune) will be
converted to float64, but complex numbers will be represented by the sum of
their real and imaginary part (in float64 form to).

5. Composed elements (arrays, channels, maps, slices, strings, structs) will be
represented by their length (or their number of fields for structs).

6. Any other element will be 0.
*/
package arithmetic
