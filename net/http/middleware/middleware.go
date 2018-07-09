package middleware

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, a ...Adapter) http.Handler {
	for i := len(a) - 1; i >= 0; i-- {
		h = a[i](h)
	}

	return h
}

func AdaptFunc(h http.HandlerFunc, a ...Adapter) http.HandlerFunc {
	return Adapt(h, a...).ServeHTTP
}
