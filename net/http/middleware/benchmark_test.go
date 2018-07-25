// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ntrrg/ntgo/net/http/middleware"
)

func benchmarkAdapt(n int, b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello, world"))

		if err != nil {
			panic(err)
		}
	})

	la := make([]middleware.Adapter, n)

	for i := 0; i < n; i++ {
		key, value := fmt.Sprintf("X-Header%d", i), "value"

		la[i] = func(h http.Handler) http.Handler {
			nh := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(key, value)
				h.ServeHTTP(w, r)
			}

			return http.HandlerFunc(nh)
		}
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)

	for i := 0; i < b.N; i++ {
		middleware.Adapt(h, la...).ServeHTTP(w, r)
	}
}

func BenchmarkAdapt(b *testing.B)      { benchmarkAdapt(1, b) }
func BenchmarkAdapt_10(b *testing.B)   { benchmarkAdapt(10, b) }
func BenchmarkAdapt_100(b *testing.B)  { benchmarkAdapt(100, b) }
func BenchmarkAdapt_1000(b *testing.B) { benchmarkAdapt(1000, b) }
