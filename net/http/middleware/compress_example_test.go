// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/ntrrg/ntgo/net/http/middleware"
)

func ExampleGzip() {
	data := []byte(`
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
`)

	h := middleware.AdaptFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write(data); err != nil {
				// Error handling
			}
		},

		middleware.Gzip(-1),
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept-Encoding", "gzip")
	h.ServeHTTP(w, r)

	// Response
	res := w.Result()
	gzdata, err := ioutil.ReadAll(res.Body)

	if err != nil {
		// Error handling
	}

	fmt.Println("Original:", len(data), "bytes")
	fmt.Println("Compressed:", len(gzdata), "bytes")
	// Output:
	// Original: 1280 bytes
	// Compressed: 637 bytes
}