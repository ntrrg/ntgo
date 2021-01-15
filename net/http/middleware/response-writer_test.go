// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.ntrrg.dev/ntgo/net/http/middleware"
)

func TestAdaptResposeWriter(t *testing.T) {
	t.Parallel()

	// All interfaces

	rw := httptest.NewRecorder()
	w := http.ResponseWriter(&adaptedRWFRF{rw})
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	_, isFlusher := w.(http.Flusher)
	_, isReaderFrom := w.(io.ReaderFrom)

	if !isFlusher || !isReaderFrom {
		t.Error("ResponseWriter should implement every other interfaces")
	}

	// http.ResponseWriter

	w = http.ResponseWriter(&adaptedRW{rw})
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	_, isFlusher = w.(http.Flusher)
	_, isReaderFrom = w.(io.ReaderFrom)

	if isReaderFrom || isFlusher {
		t.Error("ResponseWriter should only implement http.ResponseWriter")
	}

	// http.Fluser

	w = http.ResponseWriter(&adaptedRWF{rw})
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	_, isFlusher = w.(http.Flusher)
	_, isReaderFrom = w.(io.ReaderFrom)

	if !isFlusher || isReaderFrom {
		t.Error("ResponseWriter should only implement http.Flusher")
	}

	// io.ReaderFrom

	w = http.ResponseWriter(&adaptedRWRF{rw})
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	_, isFlusher = w.(http.Flusher)
	_, isReaderFrom = w.(io.ReaderFrom)

	if !isReaderFrom || isFlusher {
		t.Error("ResponseWriter should only implement io.ReaderFrom")
	}
}

func TestAdaptedRW_Unwrap(t *testing.T) {
	t.Parallel()

	w := http.ResponseWriter(httptest.NewRecorder())
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	aw, ok := w.(middleware.ResponseWriteAdapter)
	if !ok {
		t.Fatal("the ResponseWriter is not adapted")
	}

	w = aw.Unwrap()

	aw, ok = w.(middleware.ResponseWriteAdapter)
	if !ok {
		t.Fatal("the unwraped ResponseWriter should be a ResponseWriteAdapter")
	}

	w = aw.Unwrap()

	_, ok = w.(middleware.ResponseWriteAdapter)
	if ok {
		t.Fatal("the last ResponseWriter should not be a ResponseWriteAdapter")
	}
}

func TestAdaptedRW_Header(t *testing.T) {
	t.Parallel()

	w := http.ResponseWriter(httptest.NewRecorder())
	aw := middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	aw.Header().Set("X-Testing-Head", "false")

	aw = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{
		Header: func() http.Header {
			header := w.Header()
			header.Set("X-Testing-Head", "true")
			return header
		},
	})

	if aw.Header().Get("X-Testing-Head") != "true" {
		t.Error("adapted ResponseWriter is not using adapted Header")
	}
}

func TestAdaptedRW_Write(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	w := http.ResponseWriter(rw)
	aw := middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	if _, err := aw.Write([]byte("hello, world")); err != nil {
		t.Fatal(err)
	}

	aw = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{
		Write: func(buf []byte) (int, error) {
			rw.Body.Reset()
			return w.Write(append(buf, []byte(" (adapted)")...))
		},
	})

	if _, err := aw.Write([]byte("hello, world")); err != nil {
		t.Fatal(err)
	}

	if rw.Body.String() != "hello, world (adapted)" {
		t.Error("adapted ResponseWriter is not using adapted Write")
	}
}

func TestAdaptedRW_WriteHeader(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	w := http.ResponseWriter(rw)
	aw := middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	aw.WriteHeader(http.StatusBadGateway)

	aw = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{
		WriteHeader: func(_ int) {
			rw.Code = http.StatusSeeOther
		},
	})

	aw.WriteHeader(http.StatusOK)

	res := rw.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusSeeOther {
		t.Error("adapted ResponseWriter is not using adapted WriteHeader")
	}
}

func TestAdaptedRW_Flush(t *testing.T) {
	t.Parallel()

	rw := httptest.NewRecorder()
	w := http.ResponseWriter(rw)
	aw := middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	aw.(http.Flusher).Flush()

	aw = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{
		Flush: func() {
			rw.Flushed = false
		},
	})

	aw.(http.Flusher).Flush()

	if rw.Flushed == true {
		t.Error("adapted ResponseWriter is not using adapted Flush")
	}
}

func TestAdaptedRW_ReadFrom(t *testing.T) {
	t.Parallel()

	src := bytes.NewReader([]byte("hello, world"))
	rw := httptest.NewRecorder()
	w := &adaptedRWRF{rw}
	aw := middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})

	if _, err := aw.(io.ReaderFrom).ReadFrom(src); err != nil {
		t.Fatal(err)
	}

	aw = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{
		ReadFrom: func(_ io.Reader) (int64, error) {
			rw.Body.Reset()
			return w.ReadFrom(bytes.NewReader([]byte("hello, world (adapted)")))
		},
	})

	if _, err := aw.(io.ReaderFrom).ReadFrom(src); err != nil {
		t.Fatal(err)
	}

	if rw.Body.String() != "hello, world (adapted)" {
		t.Error("adapted ResponseWriter is not using adapted ReadFrom")
	}
}

func TestIsAdaptedResponseWriter(t *testing.T) {
	t.Parallel()

	w := http.ResponseWriter(httptest.NewRecorder())
	if middleware.IsAdaptedResponseWriter(w) {
		t.Fatal("IsAdaptedResponseWriter reports a false positive")
	}

	w = middleware.AdaptResponseWriter(w, middleware.ResponseWriterMethods{})
	if !middleware.IsAdaptedResponseWriter(w) {
		t.Fatal("IsAdaptedResponseWriter reports a false negative")
	}
}

type adaptedRW struct {
	http.ResponseWriter
}

type adaptedRWF struct {
	http.ResponseWriter
}

func (aw *adaptedRWF) Flush() {
	aw.ResponseWriter.(http.Flusher).Flush()
}

type adaptedRWRF struct {
	http.ResponseWriter
}

func (aw *adaptedRWRF) ReadFrom(src io.Reader) (int64, error) {
	return io.Copy(aw, src)
}

type adaptedRWFRF struct {
	http.ResponseWriter
}

func (aw *adaptedRWFRF) Flush() {
	aw.ResponseWriter.(http.Flusher).Flush()
}

func (aw *adaptedRWFRF) ReadFrom(src io.Reader) (int64, error) {
	return io.Copy(aw, src)
}
