// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package os

import (
	"os"
)

// GetenvOr retrieves the value of the environment variable named by the key.
// If the variable is empty or absent, val will be used as the return value.
func GetenvOr(key, val string) string {
	v := os.Getenv(key)
	if v == "" {
		return val
	}

	return v
}
