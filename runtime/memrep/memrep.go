// Copyright 2020 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package memrep

import (
	"reflect"
	"unsafe"
)

// Representation is a human-friendly representation generated from a Go type
// and its machine representation. Notice that Bytes is platform-dependent.
type Representation struct {
	Type  string
	Size  uintptr
	Bytes []byte
}

// Read creates a memory representation of x.
func Read(x interface{}) Representation {
	p := (*[2]uintptr)(unsafe.Pointer(&x))[1] // Get interface data pointer.
	v := reflect.ValueOf(x)
	s := v.Type().Size()

	return Representation{
		Type:  v.Type().String(),
		Size:  s,
		Bytes: readBytes(p, s),
	}
}

// ReadArray creates a memory representation of the underlying array of x.
func ReadArray(x interface{}) Representation {
	v := reflect.ValueOf(x)
	s := v.Type().Elem().Size() * uintptr(v.Len())

	return Representation{
		Type:  "array from " + v.Type().String(),
		Size:  s,
		Bytes: readBytes(v.Pointer(), s),
	}
}

func readBytes(ptr, n uintptr) []byte {
	r := make([]byte, n)

	for i := range r {
		r[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	return r
}
