// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestAddHeader(t *testing.T) {
	t.Parallel()

	cases := []struct {
		key    string
		values []string
	}{
		{key: "X-Header", values: []string{"x-value", "y-value"}},
		{key: "X-My-Header", values: []string{"x", "y", "z"}},
	}

	for i, c := range cases {
		adapters := make([]middleware.Adapter, len(c.values))

		for j, value := range c.values {
			adapters[j] = middleware.AddHeader(c.key, value)
		}

		h := middleware.AdaptFunc(
			func(w http.ResponseWriter, r *http.Request) {},
			adapters...,
		)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodHead, "/", nil)
		h.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()

		header := res.Header[c.key]

		if len(header) != len(c.values) {
			msg := "TC#%v: len(%+v) got %v, want %v"
			t.Errorf(msg, i, header, len(header), len(c.values))
		}
	}
}

func TestDelHeader(t *testing.T) {
	t.Parallel()

	key := "X-Header"

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.SetHeader(key, "Abc"),
		middleware.AddHeader(key, "Def"),
		middleware.DelHeader(key),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/", nil)
	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	if header := res.Header.Get(key); header != "" {
		t.Errorf("The 'X-Header' header stills have values: %v", header)
	}
}

func TestSetHeader(t *testing.T) {
	t.Parallel()

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.SetHeader("X-Header", "Abc"),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/", nil)
	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	if header := res.Header.Get("X-Header"); header != "Abc" {
		t.Errorf("Bad header value, got %v, want %v", res, "Abc")
	}
}
