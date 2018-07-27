// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/ntrrg/ntgo/net/http/middleware"
)

func ExampleStripPrefix() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte(r.URL.Path)); err != nil {
				// Error handling
			}
		},

		middleware.StripPrefix("/api"),
	)

	// http.Handle("/api/", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/api/", nil)
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		// Error handling
	}

	fmt.Println(string(data))
	// Output:
	// /
}
