// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ntrrg/ntgo/net/http/middleware"
)

func TestAddHeader(t *testing.T) {
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
		res := w.Result().Header[c.key]

		if len(res) != len(c.values) {
			msg := "TC#%v: len(%+v) got %v, want %v"
			t.Errorf(msg, i, res, len(res), len(c.values))
		}
	}
}

func TestCache(t *testing.T) {
	cases := []struct {
		method, value, want string
	}{
		{
			method: http.MethodGet,
			value:  "max-age=3600, s-max-age=3600",
			want:   "max-age=3600, s-max-age=3600",
		},
		{
			method: http.MethodHead,
			value:  "private, max-age=3600",
			want:   "",
		},
	}

	for i, c := range cases {
		h := middleware.AdaptFunc(
			func(w http.ResponseWriter, r *http.Request) {},
			middleware.Cache(c.value),
		)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(c.method, "/", nil)
		h.ServeHTTP(w, r)
		res := w.Result().Header.Get("Cache-Control")

		if res != c.want {
			t.Errorf("TC#%v: 'Cache-Control' == %v, want: %v", i, res, c.want)
		}
	}
}

func TestDelHeader(t *testing.T) {
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
	res := w.Result().Header.Get(key)

	if res != "" {
		t.Errorf("The 'X-Header' header stills have values -> %v", res)
	}
}

func TestJSONRequest(t *testing.T) {
	cases := []struct {
		method, ct string
		status     int
	}{
		{
			method: http.MethodGet,
			ct:     "application/json; charset=utf-8",
			status: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodPost,
			ct:     "application/json; charset=utf-8",
			status: http.StatusOK,
		},
		{
			method: http.MethodPut,
			ct:     "text/plain; charset=utf-8",
			status: http.StatusUnsupportedMediaType,
		},
		{
			method: http.MethodPatch,
			ct:     "application/json",
			status: http.StatusOK,
		},
		{
			method: http.MethodDelete,
			ct:     "text/plain; charset=utf-8",
			status: http.StatusMethodNotAllowed,
		},
	}

	for i, c := range cases {
		h := middleware.AdaptFunc(
			func(w http.ResponseWriter, r *http.Request) {},
			middleware.JSONRequest(""),
		)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(c.method, "/", nil)
		r.Header.Set("Content-Type", c.ct)
		h.ServeHTTP(w, r)
		res := w.Result().StatusCode

		if res != c.status {
			t.Errorf("TC#%v: got %v, want %v", i, res, c.status)
		}
	}
}

func TestJSONResponse(t *testing.T) {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.JSONResponse(),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(w, r)
	res := w.Result().Header.Get("Content-Type")

	if res != "application/json; charset=utf-8" {
		t.Errorf("Bad header value -> %v", res)
	}
}

func TestSetHeader(t *testing.T) {
	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.SetHeader("X-Header", "Abc"),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/", nil)
	h.ServeHTTP(w, r)
	res := w.Result().Header.Get("X-Header")

	if res != "Abc" {
		t.Errorf("Bad header value, got %v, want %v", res, "Abc")
	}
}
