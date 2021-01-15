// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestReplace(t *testing.T) {
	t.Parallel()

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte(r.URL.Path)); err != nil {
				t.Error(err)
			}
		},

		middleware.Replace("pi", "314"),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/api/", nil)
	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "/a314/" {
		t.Errorf("Got %v, want /a314/", string(data))
	}
}

func TestStripPrefix(t *testing.T) {
	t.Parallel()

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte(r.URL.Path)); err != nil {
				t.Error(err)
			}
		},

		middleware.StripPrefix("/api"),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodHead, "/api/", nil)
	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "/" {
		t.Errorf("Got %v, want /", string(data))
	}
}
