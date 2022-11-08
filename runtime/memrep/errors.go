// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package memrep

import (
	ntruntime "go.ntrrg.dev/ntgo/runtime"
)

// Err is the main error group for this package.
var Err = ntruntime.Err.New("memrep", "memrep package errors")
