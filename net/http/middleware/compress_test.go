// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestGzip(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want int
	}{
		{
			in: `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas luctus ex non
lobortis sagittis. Nulla feugiat velit a eros ultrices mollis. Quisque pulvinar
nec odio eu gravida. Etiam at gravida nisl. Donec consequat eget libero nec
consectetur. Fusce posuere a mauris nec pellentesque. Ut quis sem at mi dictum
feugiat vitae id est. Vestibulum ante ipsum primis in faucibus orci luctus et
ultrices posuere cubilia Curae; Etiam ultrices tempus ex ut sagittis. Nunc
felis dui, varius at tincidunt nec, consectetur ac turpis. Mauris id porttitor
nulla. Suspendisse euismod urna eget nunc interdum ornare. Donec luctus augue
elementum congue commodo.

Nullam vel efficitur dui. Quisque dignissim mauris non mi imperdiet, aliquet
faucibus velit auctor. Vivamus felis quam, cursus quis ultrices faucibus,
lobortis non nisi. Proin nec pellentesque nulla. Vivamus ut eros accumsan,
egestas nulla a, aliquet nulla. Suspendisse potenti. Nunc sollicitudin sapien
dolor, quis sagittis risus ultricies vitae. Nam non magna lacinia urna vehicula
mattis ut a nunc. Integer tincidunt urna commodo justo lobortis porttitor.
Vestibulum aliquet consequat magna at molestie. Sed tortor nulla, hendrerit sed
imperdiet sit amet, viverra ac nisl. Quisque porttitor vestibulum massa eget
consequat.
`,

			want: 637,
		},
	}

	for i, c := range cases {
		i, c := i, c

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Set("Accept-Encoding", "gzip")

		h := middleware.AdaptFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if _, err := w.Write([]byte(c.in)); err != nil {
					t.Fatal(err)
				}
			},

			middleware.Gzip(-1),
		)

		h.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()

		gzdata, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		if len(gzdata) != c.want {
			t.Errorf("TC#%v: got %v, want %v", i, len(gzdata), c.want)
		}
	}
}

func TestGzip_adaptedResponseWriter(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	w := http.ResponseWriter(&adaptedRWFRF{rw})
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept-Encoding", "gzip")

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("hello, world")); err != nil {
				t.Fatal(err)
			}

			f, ok := w.(http.Flusher)
			if !ok {
				t.Fatal("ResposeWriter should implement http.Flusher")
			}

			f.Flush()

			rf, ok := w.(io.ReaderFrom)
			if !ok {
				t.Fatal("ResposeWriter should implement io.ReaderFrom")
			}

			src := bytes.NewReader([]byte("\ngoodbye, world"))
			_, err := rf.ReadFrom(src)
			if err != nil {
				t.Error(err)
			}
		},
		middleware.Gzip(-1),
	)

	h.ServeHTTP(w, r)

	res := rw.Result()
	defer res.Body.Close()

	gzr, err := gzip.NewReader(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	defer gzr.Close()

	data, err := io.ReadAll(gzr)
	if err != nil {
		t.Fatal(err)
	}

	got := string(data)
	want := "hello, world\ngoodbye, world"

	if got != want {
		t.Errorf("invalid response payload.\ngot: %s\nwant: %s", got, want)
	}
}

func TestGzip_noHeader(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {},
		middleware.Gzip(-1),
	)

	h.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.Header.Get("Content-Encoding") == "gzip" {
		t.Error("the response should not be compressed")
	}
}
