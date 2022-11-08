// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package middleware

import (
	"io"
	"net/http"
)

// AdaptResponseWriter wraps a ResponseWriter with custom methods. This allows
// to overwrite only specific methods without compromising other interface
// implementations.
//
// This technique is based on https://github.com/felixge/httpsnoop.
func AdaptResponseWriter(
	w http.ResponseWriter,
	m ResponseWriterMethods,
) http.ResponseWriter {
	_, isFlusher := w.(http.Flusher)
	_, isReaderFrom := w.(io.ReaderFrom)
	aw := &adaptedRW{w, m}

	switch {
	case isFlusher && isReaderFrom:
		return struct {
			ResponseWriteAdapter
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{aw, aw, aw, aw}
	case isFlusher:
		return struct {
			ResponseWriteAdapter
			http.ResponseWriter
			http.Flusher
		}{aw, aw, aw}
	case isReaderFrom:
		return struct {
			ResponseWriteAdapter
			http.ResponseWriter
			io.ReaderFrom
		}{aw, aw, aw}
	}

	return struct {
		ResponseWriteAdapter
		http.ResponseWriter
	}{aw, aw}
}

type adaptedRW struct {
	w http.ResponseWriter
	m ResponseWriterMethods
}

func (aw *adaptedRW) Unwrap() http.ResponseWriter {
	return aw.w
}

// http.ResponseWriter implementation

func (aw *adaptedRW) Header() http.Header {
	if aw.m.Header != nil {
		return aw.m.Header()
	}

	return aw.w.Header()
}

func (aw *adaptedRW) Write(p []byte) (int, error) {
	if aw.m.Write != nil {
		return aw.m.Write(p)
	}

	return aw.w.Write(p) //nolint:wrapcheck
}

func (aw *adaptedRW) WriteHeader(statusCode int) {
	if aw.m.WriteHeader != nil {
		aw.m.WriteHeader(statusCode)
		return
	}

	aw.w.WriteHeader(statusCode)
}

// http.Flusher implementation

func (aw *adaptedRW) Flush() {
	if aw.m.Flush != nil {
		aw.m.Flush()
		return
	}

	f := aw.w.(http.Flusher) //nolint:errcheck,forcetypeassert
	f.Flush()
}

// io.ReaderFrom implementation

func (aw *adaptedRW) ReadFrom(r io.Reader) (int64, error) {
	if aw.m.ReadFrom != nil {
		return aw.m.ReadFrom(r)
	}

	rf := aw.w.(io.ReaderFrom) //nolint:errcheck,forcetypeassert

	return rf.ReadFrom(r) //nolint:wrapcheck
}

// IsAdaptedResponseWriter reports if the given http.ResponseWriter were
// adapted previously.
func IsAdaptedResponseWriter(w http.ResponseWriter) bool {
	_, ok := w.(ResponseWriteAdapter)
	return ok
}

// A ResponseWriteAdapter represents an adapted http.ResponseWriter.
type ResponseWriteAdapter interface {
	// Unwrap allows access to the underlying http.ResponseWriter.
	Unwrap() http.ResponseWriter
}

// ResponseWriterMethods is a set of methods used to wrap a http.ReponseWriter.
type ResponseWriterMethods struct {
	// http.ResponseWriter
	Header      func() http.Header
	Write       func([]byte) (int, error)
	WriteHeader func(int)

	// http.Flusher
	Flush func()

	// io.ReaderFrom
	ReadFrom func(io.Reader) (int64, error)
}
