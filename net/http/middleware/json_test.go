// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestJSONRequest(t *testing.T) {
	t.Parallel()

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
			middleware.JSONRequest(),
		)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(c.method, "/", nil)
		r.Header.Set("Content-Type", c.ct)
		h.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()

		status := res.StatusCode
		if status != c.status {
			t.Errorf("TC#%v: got %v, want %v", i, status, c.status)
		}
	}
}

func TestJSONResponse(t *testing.T) {
	t.Parallel()

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.JSONResponse(),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	header := res.Header.Get("Content-Type")
	if header != "application/json; charset=utf-8" {
		t.Errorf("Bad header value: %v", res)
	}
}
