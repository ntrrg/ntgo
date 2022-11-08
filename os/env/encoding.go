// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package env

// Decoder is a plain text decoder for a specific type.
type Decoder[T any] func(string) (T, error)

func String(v string) (string, error) {
	return v, nil
}
