// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func Example() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	nthttp.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data := []byte(`{ "msg": "hello, world" }`)
			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

			if _, err := w.Write(data); err != nil {
				// Error handling
			}
		},

		nthttp.SetHeader("Content-Type", "application/json; charset=utf-8"),
	).ServeHTTP(w, r)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		// Error handling
	}

	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println(res.Header.Get("Content-Length"))
	fmt.Println(string(data))
	// Output:
	// application/json; charset=utf-8
	// 25
	// { "msg": "hello, world" }
}
