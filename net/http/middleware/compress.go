// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Gzip compresses the response body. The compression level is given as an
// integer value according to the compress/flate package.
func Gzip(level int) Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			wh := w.Header()

			// Prevent proxy caches corruption
			wh.Add("Vary", "Accept-Encoding")

			// Ignore requests from clients that don't support/want GZIP
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				h.ServeHTTP(w, r)
				return
			}

			gz, err := gzip.NewWriterLevel(w, level)
			if err != nil {
				panic(fmt.Errorf("bad compression level: %w", err))
			}

			defer gz.Close()
			wh.Set("Content-Encoding", "gzip")

			wh["Content-Length"] = nil

			aw := AdaptResponseWriter(w, ResponseWriterMethods{
				Write: func(buf []byte) (int, error) {
					return gz.Write(buf)
				},

				Flush: func() {
					gz.Flush()
				},

				ReadFrom: func(src io.Reader) (int64, error) {
					return io.Copy(gz, src)
				},
			})

			h.ServeHTTP(aw, r)
		}

		return http.HandlerFunc(nh)
	}
}
