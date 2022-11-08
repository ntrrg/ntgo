// Copyright 2022 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package env

import (
	"errors"
	"syscall"
)

var (
	ErrGet = Err.New("get", "cannot get environment variable value")

	ErrCannotDecode = ErrGet.New("decode", "cannot decode value")
	ErrUndefined    = ErrGet.New("undefined", "variable not defined")
)

// Must is a helper that wraps a call to a function returning a value from an
// environment variable and panics if there is any error.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// Get retrieves the value of the environment variable k. If the variable is
// absent, an empty string will be used as input for fn.
//
// See Lookup for detecting unset environment variables.
func Get[T any](k string, fn Decoder[T]) (T, error) {
	v, _ := syscall.Getenv(k)
	return decode(v, fn)
}

// GetOr retrieves the value of the environment variable k. If the variable is
// empty or absent, v will be returned as default value.
//
// See LookupOr for using v only if the environment variable is unset.
func GetOr[T any](k string, v T, fn Decoder[T]) (T, error) {
	_v, _ := syscall.Getenv(k)
	if _v == "" {
		return v, nil
	}

	return decode(_v, fn)
}

// Lookup retrieves the value of the environment variable k. If the variable is
// absent, ErrUndefined will be returned.
//
// See Get for ignoring unset environment variables.
func Lookup[T any](k string, fn Decoder[T]) (T, error) {
	var v T

	_v, ok := syscall.Getenv(k)
	if !ok {
		return v, ErrUndefined.Wrap(errors.New("'" + k + "' not found"))
	}

	return decode(_v, fn)
}

// LookupOr retrieves the value of the environment variable k. If the variable
// is absent, v will be returned as default value.
//
// See GetOr for also using v when the environment variable is empty.
func LookupOr[T any](k string, v T, fn Decoder[T]) (T, error) {
	_v, ok := syscall.Getenv(k)
	if !ok {
		return v, nil
	}

	return decode(_v, fn)
}

func decode[T any](val string, fn Decoder[T]) (v T, err error) {
	v, err = fn(val)
	if err != nil {
		return v, ErrCannotDecode.Wrap(err)
	}

	return v, nil
}
