// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import "net/http"

// Adapter is a wrapper for http.Handler. Takes an http.Handler as argument and
// creates a new one that may run code before and/or after calling the given
// handler.
type Adapter func(http.Handler) http.Handler

// Adapt wraps an http.Handler into a list of Adapters. Adapters will be
// wrapped right to left (they will be executed left to right).
func Adapt(h http.Handler, a ...Adapter) http.Handler {
	for i := len(a) - 1; i >= 0; i-- {
		h = a[i](h)
	}

	return h
}

// AdaptFunc works as Adapt but for http.HandlerFunc.
func AdaptFunc(
	h func(w http.ResponseWriter, r *http.Request),
	a ...Adapter,
) http.Handler {
	return Adapt(http.HandlerFunc(h), a...)
}
