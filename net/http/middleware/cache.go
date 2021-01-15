// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import (
	"net/http"
)

// Cache sets HTTP cache headers for GET requests.
//
// Supported directives:
//
// * public/private: whether the cached response is for any or a specific user.
//
// * max-age=TIME: cache life time in seconds. The maximum value is 1 year.
//
// * s-max-age=TIME: same as max-age, but this one has effect in proxies.
//
// * must-revalidate: force expired cached response revalidation, even in
// special circumstances (like slow connections, were cached responses are used
// even after they had expired).
//
// * proxy-revalidate: same as must-revalidate, but this one has effect in
// proxies.
//
// * no-cache: disables cache.
//
// * no-store: disables cache, even in proxies.
func Cache(directives string) Adapter {
	return func(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				w.Header().Set("Cache-Control", directives)
			}

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(nh)
	}
}
