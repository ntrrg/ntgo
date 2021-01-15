// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestCache(t *testing.T) {
	t.Parallel()

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

		res := w.Result()
		defer res.Body.Close()

		header := res.Header.Get("Cache-Control")
		if header != c.want {
			t.Errorf("TC#%v: 'Cache-Control' == %v, want: %v", i, header, c.want)
		}
	}
}
