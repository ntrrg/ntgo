package http_test

import (
	"fmt"
	"net/http"
	"testing"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func benchmarkChainHandlers(n int, b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	la := make([]http.Handler, n)

	for i := 0; i < n; i++ {
		la[i] = nthttp.SetHeader(fmt.Sprintf("X-Header%d", i), "value")
	}

	la = append(la, h)

	for i := 0; i < b.N; i++ {
		nthttp.ChainHandlers(la...)
	}
}

func BenchmarkChainHandlers(b *testing.B)      { benchmarkChainHandlers(1, b) }
func BenchmarkChainHandlers_10(b *testing.B)   { benchmarkChainHandlers(10, b) }
func BenchmarkChainHandlers_100(b *testing.B)  { benchmarkChainHandlers(100, b) }
func BenchmarkChainHandlers_1000(b *testing.B) { benchmarkChainHandlers(1000, b) }

// Adapter pattern

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, a ...Adapter) http.Handler {
	for i := len(a) - 1; i >= 0; i-- {
		h = a[i](h)
	}

	return h
}

func SetHeaderAdapter(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		hf := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(hf)
	}
}

func benchmarkAdapter(n int, b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})

	la := make([]Adapter, n)

	for i := 0; i < n; i++ {
		la[i] = SetHeaderAdapter(fmt.Sprintf("X-Header%d", i), "value")
	}

	for i := 0; i < b.N; i++ {
		Adapt(h, la...)
	}
}

func BenchmarkAdapter(b *testing.B)      { benchmarkAdapter(1, b) }
func BenchmarkAdapter_10(b *testing.B)   { benchmarkAdapter(10, b) }
func BenchmarkAdapter_100(b *testing.B)  { benchmarkAdapter(100, b) }
func BenchmarkAdapter_1000(b *testing.B) { benchmarkAdapter(1000, b) }
