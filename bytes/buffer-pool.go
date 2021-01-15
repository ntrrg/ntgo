// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes

import (
	"bytes"
)

// BufferPool is a pool implementation of bytes.Buffer. This allows to reuse
// buffers, which means less allocations. Buffers are guaranteed to live in
// memory, which is a key difference between sync.Pool.
type BufferPool struct {
	pool chan *bytes.Buffer
	max  int
}

// NewBufferPool creates a pool that will hold up to n buffers. Any new buffer
// will be created with up to max bytes.
func NewBufferPool(n, max int) BufferPool {
	return BufferPool{
		pool: make(chan *bytes.Buffer, n),
		max:  max,
	}
}

// Add appends buf to the pool. Add behaves like AddWait, but if there is no
// more room for buffers, buf will be discarded.
func (bp *BufferPool) Add(buf *bytes.Buffer) {
	if bp.Len() == bp.Cap() {
		return
	}

	bp.AddWait(buf)
}

// AddWait appends buf to the pool. If there is no room for more buffers,
// AddWait will wait until buf can be appended. If buf is nil or has more than
// bp.max bytes, it will be discarded and a new buffer will be allocated. After
// calling AddWait, buf must not be used anymore.
func (bp *BufferPool) AddWait(buf *bytes.Buffer) {
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

// Clear discards all the buffers in the pool.
func (bp *BufferPool) Clear() {
	for len(bp.pool) > 0 {
		<-bp.pool
	}
}

// Fill fills the pool up to its available capacity.
func (bp *BufferPool) Fill() {
	n := bp.Cap() - bp.Len()

	for i := 0; i < n; i++ {
		bp.Add(nil)
	}
}

// Get returns a buffer from the pool. If there are no buffers in the pool, a
// new one with bp.max bytes will be returned.
func (bp *BufferPool) Get() (buf *bytes.Buffer) {
	select {
	case buf = <-bp.pool:
	default:
		buf = newBufferL(bp.max)
	}

	return buf
}

// GetWait returns a buffer from the pool. If there are no buffers in the pool,
// this will wait until one is available.
func (bp *BufferPool) GetWait() *bytes.Buffer {
	return <-bp.pool
}

// Len returns the amount of buffers in the pool.
func (bp *BufferPool) Len() int {
	return len(bp.pool)
}

func newBufferL(l int) *bytes.Buffer {
	s := make([]byte, 0, l)

	return bytes.NewBuffer(s)
}
