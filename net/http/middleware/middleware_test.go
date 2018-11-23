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

func ExampleAdapt() {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.JSONResponse(),
		middleware.Cache("max-age=3600, s-max-age=3600"),
	)

	// http.Handle("/", h)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(w, r)
	res := w.Result()
	fmt.Printf("Status: %v\n", res.Status)
	fmt.Printf("Content-Type: %v\n", res.Header["Content-Type"][0])
	fmt.Printf("Cache-Control: %+v", res.Header["Cache-Control"][0])
	// Output:
	// Status: 200 OK
	// Content-Type: application/json; charset=utf-8
	// Cache-Control: max-age=3600, s-max-age=3600
}

type headersIn []struct {
	key, value string
}

type headersWant map[string]string

func testAdapt(f interface{}, t *testing.T) {
	cases := []struct {
		in   headersIn
		want headersWant
	}{
		{
			in: headersIn{
				{"X-Header", "x-value"},
				{"Content-Type", "application/json; charset=utf-8"},
				{"X-Header", "y-value"},
			},

			want: headersWant{
				"X-Header":     "y-value",
				"Content-Type": "application/json; charset=utf-8",
			},
		},
	}

	for i, c := range cases {
		adapters := make([]middleware.Adapter, len(c.in))

		for j, h := range c.in {
			key, value := h.key, h.value

			adapters[j] = func(h http.Handler) http.Handler {
				nh := func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set(key, value)
					h.ServeHTTP(w, r)
				}

				return http.HandlerFunc(nh)
			}
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodHead, "/", nil)

		switch v := f.(type) {
		case func(http.Handler, ...middleware.Adapter) http.Handler:
			v(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				adapters...,
			).ServeHTTP(w, r)

		case func(func(http.ResponseWriter, *http.Request), ...middleware.Adapter) http.Handler:
			v(
				func(w http.ResponseWriter, r *http.Request) {},
				adapters...,
			).ServeHTTP(w, r)
		}

		headers := w.Header()

		for key, value := range c.want {
			if value != headers[key][0] {
				t.Errorf("TC#%v: got %q, want %q", i, headers[key], value)
				break
			}
		}
	}
}

func TestAdapt(t *testing.T)     { testAdapt(middleware.Adapt, t) }
func TestAdaptFunc(t *testing.T) { testAdapt(middleware.AdaptFunc, t) }

func benchmarkAdapt(n int, b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello, world"))

		if err != nil {
			panic(err)
		}
	})

	adapters := make([]middleware.Adapter, n)

	for i := 0; i < n; i++ {
		key, value := fmt.Sprintf("X-Header%d", i), "value"

		adapters[i] = func(h http.Handler) http.Handler {
			nh := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(key, value)
				h.ServeHTTP(w, r)
			}

			return http.HandlerFunc(nh)
		}
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/", nil)

	for i := 0; i < b.N; i++ {
		middleware.Adapt(h, adapters...).ServeHTTP(w, r)
	}
}

func BenchmarkAdapt(b *testing.B)      { benchmarkAdapt(1, b) }
func BenchmarkAdapt_10(b *testing.B)   { benchmarkAdapt(10, b) }
func BenchmarkAdapt_100(b *testing.B)  { benchmarkAdapt(100, b) }
func BenchmarkAdapt_1000(b *testing.B) { benchmarkAdapt(1000, b) }
