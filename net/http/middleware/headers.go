// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import (
	"net/http"
)

// AddHeader creates/appends a HTTP header before calling the http.Handler.
func AddHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(nh)
	}
}

// DelHeader removes a HTTP header before calling the http.Handler.
func DelHeader(key string) Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Del(key)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(nh)
	}
}

// SetHeader creates/replaces a HTTP header before calling the http.Handler.
func SetHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(nh)
	}
}
