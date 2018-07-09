package middleware

import "net/http"

func SetHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		hf := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(hf)
	}
}
