// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

/*
Package middleware provides flexibility at the HTTP request/response process.
For this purpose, the Adapter pattern is used, which consist in wrapping
handlers. An adapter may run code before and/or after the handler it is
wrapping.

	func MyAdapter(h http.Handler) http.Handler {
		nh := func(w http.ResponseWriter, r *http.Request) {
			// Code that run before
			h.ServeHTTP(w, r) // Breaks the flow if not used
			// Code that run after
		}

		return http.HandlerFunc(nh)
	}

When multiple adapters are used, the result is a sequence of wrappers and its
execution flow depends in the order that adapters were given.

	Adapt(h, f1, f2, f3)

1. f1 before code

2. f2 before code

3. f3 before code

4. h

5. f3 after code

6. f2 after code

7. f1 after code

Some adapters my require to change the behavior of the ResponseWriter. Since
the underlying type of a ResponseWriter from the stdlib implements more than
this interface (which is a bad design decision), simply wrapping it with a
custom type will hide other interface implementations. This leads to several
side effects during request processing and makes some middleware completely
unusable.

To solve this, AdaptResponseWriter must be used. Supported interfaces are:

* http.Flusher

* io.ReaderFrom

Unsupported interfaces are:

* http.CloseNotifier: this was deprecated in favor of Request.Context.

* http.Hijacker: I don't like Websockets (or the way they are implemented) and
hijacking HTTP is a hacky workaround (also, ResponseWriters for HTTP versions
higher than 1.1 don't implement this interface).

* http.Pusher: web browsers are deprecating HTTP/2 Server Push.
*/
package middleware

// API Status: testing
