// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package runtime

import (
	"unsafe"
)

// IsLittleEndian checks if the current platform uses big-endian.
func IsBigEndian() bool {
	return !IsLittleEndian()
}

// IsLittleEndian checks if the current platform uses little-endian.
func IsLittleEndian() bool {
	var x int16 = 0x0011
	return *(*byte)(unsafe.Pointer(&x)) == 0x11
}

// WordSize returns the current platform word size.
func WordSize() uintptr {
	return unsafe.Sizeof(uintptr(0))
}
