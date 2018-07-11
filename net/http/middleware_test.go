package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

type HTTPHeader struct {
	key, value string
}

func testChain(f interface{}, t *testing.T) {
	headers := []HTTPHeader{
		{"X-Header", "x-value"},
		{"Content-Type", "application/json; charset=utf-8"},
		{"X-Header", "y-value"},
	}

	results := map[string][]string{
		"X-Header":     {"y-value"},
		"Content-Type": {"application/json; charset=utf-8"},
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("HEAD", "/", nil)

	switch v := f.(type) {
	case func(...http.Handler) http.Handler:
		middleware := make([]http.Handler, len(headers))

		for i, h := range headers {
			middleware[i] = nthttp.SetHeader(h.key, h.value)
		}

		v(middleware...).ServeHTTP(w, r)
	case func(...http.HandlerFunc) http.HandlerFunc:
		middleware := make([]http.HandlerFunc, len(headers))

		for i, h := range headers {
			middleware[i] = nthttp.SetHeader(h.key, h.value).ServeHTTP
		}

		v(middleware...).ServeHTTP(w, r)
	default:
		t.Fatalf("Wrong function type %T", f)
	}

	rHeaders := w.Header()

	if len(rHeaders) != len(results) {
		t.Errorf(
			"Got %d header(s); want = %v\n%v",
			len(rHeaders),
			len(results),
			rHeaders,
		)
	}

	for k, got := range rHeaders {
		want := results[k]

		for i, v := range want {
			if v != got[i] {
				t.Errorf("%v == %q, want: %q", k, got, want)
				break
			}
		}
	}
}

func TestChainHandlers(t *testing.T)     { testChain(nthttp.ChainHandlers, t) }
func TestChainHandlerFuncs(t *testing.T) { testChain(nthttp.ChainHandlerFuncs, t) }
