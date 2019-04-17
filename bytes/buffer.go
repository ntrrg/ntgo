// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes

import (
	"bytes"
)

// BufferPool is a pool implementation of bytes.Buffer. This allows to reuse
// buffers, which means less allocations.
type BufferPool struct {
	pool chan *bytes.Buffer
	max  int
}

// NewBufferPool creates a pool that will hold up to n buffers. Any buffer will
// have up to max bytes.
func NewBufferPool(n, max int) *BufferPool {
	return &BufferPool{
		pool: make(chan *bytes.Buffer, n),
		max:  max,
	}
}

// Add appends the given buffer to the pool. If buf is nil or has more than
// bp.max bytes, a new buffer will be used.
func (bp *BufferPool) Add(buf *bytes.Buffer) {
	if buf == nil || buf.Len() > bp.max {
		buf = newBufferL(bp.max)
	} else {
		buf.Reset()
	}

	bp.pool <- buf
}

// Cap returns the pool capacity.
func (bp *BufferPool) Cap() int {
	return cap(bp.pool)
}

// Clear discards all the buffers in then pool.
func (bp *BufferPool) Clear() {
clearing:
	for {
		select {
		case <-bp.pool:
		default:
			break clearing
		}
	}
}

// Fill fills the pool up to its available capacity.
func (bp *BufferPool) Fill() {
	n := bp.Cap() - bp.Len()
	for i := 0; i < n; i++ {
		bp.Add(nil)
	}
}

// Get returns a buffer from the pool. If there are no buffers in the pool a
// new one with bp.max bytes will be returned.
func (bp *BufferPool) Get() (buf *bytes.Buffer) {
	select {
	case buf = <-bp.pool:
	default:
		buf = newBufferL(bp.max)
	}

	return buf
}

// Len returns the amount of buffers in the pool.
func (bp *BufferPool) Len() int {
	return len(bp.pool)
}

func newBufferL(l int) *bytes.Buffer {
	s := make([]byte, 0, l)
	return bytes.NewBuffer(s)
}
