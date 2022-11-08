// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package env

import (
	ntos "go.ntrrg.dev/ntgo/os"
)

// Err is the main error group for this package.
var Err = ntos.Err.New("env", "env package errors")
