package http

import "net/http"

// ChainHandlers chains a list of http.Handlers.
func ChainHandlers(handlers ...http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			h.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(h)
}

// ChainHandlerFuncs chains a list of http.HandlersFuncs.
func ChainHandlerFuncs(handlers ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			h(w, r)
		}
	}
}

// SetHeader creates/replace a HTTP header.
func SetHeader(key, value string) http.Handler {
	hf := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(key, value)
	}

	return http.HandlerFunc(hf)
}
