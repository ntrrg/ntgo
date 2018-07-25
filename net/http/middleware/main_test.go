// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ntrrg/ntgo/net/http/middleware"
)

func testAdapt(f interface{}, t *testing.T) {
	in := []struct {
		key, value string
	}{
		{"X-Header", "x-value"},
		{"Content-Type", "application/json; charset=utf-8"},
		{"X-Header", "y-value"},
	}

	want := map[string][]string{
		"X-Header":     {"y-value"},
		"Content-Type": {"application/json; charset=utf-8"},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)

	al := make([]middleware.Adapter, len(in))

	for i, h := range in {
		key, value := h.key, h.value

		al[i] = func(h http.Handler) http.Handler {
			nh := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(key, value)
				h.ServeHTTP(w, r)
			}

			return http.HandlerFunc(nh)
		}
	}

	switch v := f.(type) {
	case func(http.Handler, ...middleware.Adapter) http.Handler:
		v(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			al...,
		).ServeHTTP(w, r)
	case func(http.HandlerFunc, ...middleware.Adapter) http.Handler:
		v(
			func(w http.ResponseWriter, r *http.Request) {},
			al...,
		).ServeHTTP(w, r)
	default:
		t.Errorf("Wrong function type %T", f)
	}

	headers := w.Header()

	if len(headers) != len(want) {
		t.Errorf(
			"Got %d header(s); want = %v\n%v",
			len(headers),
			len(want),
			headers,
		)
	}

	for k, got := range headers {
		for i, v := range want[k] {
			if v != got[i] {
				t.Errorf("%v == %q, want: %q", k, got, want[k])
				break
			}
		}
	}
}

func TestAdapt(t *testing.T)     { testAdapt(middleware.Adapt, t) }
func TestAdaptFunc(t *testing.T) { testAdapt(middleware.AdaptFunc, t) }
