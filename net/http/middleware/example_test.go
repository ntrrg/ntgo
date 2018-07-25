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

func ExampleAddHeader() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.AddHeader("X-My-Header", "Abc"),
		middleware.AddHeader("X-My-Header", "Def"),
	)

	// http.Handle("/", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	fmt.Println(res.Header["X-My-Header"])
	// Output:
	// [Abc Def]
}

func ExampleDelHeader() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.SetHeader("X-My-Header", "Abc"),
		middleware.DelHeader("X-My-Header"),
	)

	// http.Handle("/", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	fmt.Println(res.Header.Get("X-My-Header"))
	// Output:
}

func ExampleJSONResponse() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data := []byte(`{ "msg": "hello, world" }`)
			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

			if _, err := w.Write(data); err != nil {
				// Error handling
			}
		},

		middleware.JSONResponse(),
	)

	// http.Handle("/", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	defer res.Body.Close()
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

func ExampleSetHeader() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-My-Header", "Xyz")
		},

		middleware.SetHeader("X-My-Header", "Abc"),
		middleware.SetHeader("X-My-Second-Header", "Def"),
	)

	// http.Handle("/", h)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	fmt.Println(res.Header["X-My-Header"])
	fmt.Println(res.Header.Get("X-My-Second-Header"))
	// Output:
	// [Xyz]
	// Def
}
