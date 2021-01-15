// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import (
	"net/http"
	"strings"
)

// JSONRequest checks that request has the appropriate HTTP method and the
// appropriate 'Content-Type' header. Responds with http.StatusMethodNotAllowed
// if the used method is not one of POST, PUT or PATCH. Responds with
// http.StatusUnsupportedMediaType if the 'Content-Type' header is not valid.
func JSONRequest() Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			methods := []string{http.MethodPost, http.MethodPut, http.MethodPatch}

			if !isAllowedMethod(r.Method, methods) {
				http.Error(w, "", http.StatusMethodNotAllowed)
				return
			}

			if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				http.Error(w, "", http.StatusUnsupportedMediaType)
				return
			}

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(nh)
	}
}

// JSONResponse prepares the response to be a JSON response.
func JSONResponse() Adapter {
	return SetHeader("Content-Type", "application/json; charset=utf-8")
}

func isAllowedMethod(m string, methods []string) bool {
	for _, x := range methods {
		if m == x {
			return true
		}
	}

	return false
}
