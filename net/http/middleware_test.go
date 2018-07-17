// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

type HTTPHeader struct {
	key, value string
}

func testAdapt(f interface{}, t *testing.T) {
	headers := []HTTPHeader{
		{"X-Header", "x-value"},
		{"Content-Type", "application/json; charset=utf-8"},
		{"X-Header", "y-value"},
	}

	results := map[string][]string{
		"X-Header":     {"y-value"},
		"Content-Type": {"application/json; charset=utf-8"},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)

	middleware := make([]nthttp.Adapter, len(headers))

	for i, h := range headers {
		middleware[i] = nthttp.SetHeader(h.key, h.value)
	}

	switch v := f.(type) {
	case func(http.Handler, ...nthttp.Adapter) http.Handler:
		v(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			middleware...,
		).ServeHTTP(w, r)
	case func(http.HandlerFunc, ...nthttp.Adapter) http.Handler:
		v(
			func(w http.ResponseWriter, r *http.Request) {},
			middleware...,
		).ServeHTTP(w, r)
	default:
		t.Fatalf("Wrong function type %T", f)
	}

	rHeaders := w.Header()

	if len(rHeaders) != len(results) {
		t.Errorf(
			"Got %d header(s); want = %v\n%v",
			len(rHeaders),
			len(results),
			rHeaders,
		)
	}

	for k, got := range rHeaders {
		want := results[k]

		for i, v := range want {
			if v != got[i] {
				t.Errorf("%v == %q, want: %q", k, got, want)
				break
			}
		}
	}
}

func TestAdapt(t *testing.T)     { testAdapt(nthttp.Adapt, t) }
func TestAdaptFunc(t *testing.T) { testAdapt(nthttp.AdaptFunc, t) }
