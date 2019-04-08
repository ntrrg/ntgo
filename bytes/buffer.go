// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes

import (
	"bytes"
)

type BufferPool struct {
	pool chan *bytes.Buffer
	max  int
}

func NewBufferPool(n, max int) *BufferPool {
	return &BufferPool{
		pool: make(chan *bytes.Buffer, n),
		max:  max,
	}
}

func (bp *BufferPool) Clear() {
	for range bp.pool {
	}
}

func (bp *BufferPool) Fill() {
	l := len(bp.pool)
	for i := 0; i < l; i++ {
		bp.pool <- newBuffer(bp.max)
	}
}

func (bp *BufferPool) Get() (buf *bytes.Buffer) {
	select {
	case buf = <-bp.pool:
	default:
		buf = newBuffer(bp.max)
	}

	return buf
}

func (bp *BufferPool) Push(buf *bytes.Buffer) {
	if buf.Len() > bp.max {
		buf = newBuffer(bp.max)
	}

	bp.pool <- buf
}

func newBuffer(l int) *bytes.Buffer {
	s := make([]byte, l)
	return bytes.NewBuffer(s)
}
