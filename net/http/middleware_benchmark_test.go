package http_test

import (
	"fmt"
	"net/http"
	"testing"

	nthttp "github.com/ntrrg/ntgo/net/http"
)

func benchmarkChainHandlers_creation(n int, b *testing.B) {
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

func BenchmarkChainHandlers_creation(b *testing.B)      { benchmarkChainHandlers_creation(1, b) }
func BenchmarkChainHandlers_creation_10(b *testing.B)   { benchmarkChainHandlers_creation(10, b) }
func BenchmarkChainHandlers_creation_100(b *testing.B)  { benchmarkChainHandlers_creation(100, b) }
func BenchmarkChainHandlers_creation_1000(b *testing.B) { benchmarkChainHandlers_creation(1000, b) }

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

func benchmarkAdapter_creation(n int, b *testing.B) {
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

func BenchmarkAdapter_creation(b *testing.B)      { benchmarkAdapter_creation(1, b) }
func BenchmarkAdapter_creation_10(b *testing.B)   { benchmarkAdapter_creation(10, b) }
func BenchmarkAdapter_creation_100(b *testing.B)  { benchmarkAdapter_creation(100, b) }
func BenchmarkAdapter_creation_1000(b *testing.B) { benchmarkAdapter_creation(1000, b) }
