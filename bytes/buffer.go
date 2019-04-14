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

func NewBufferPool(size, max int) *BufferPool {
	return &BufferPool{
		pool: make(chan *bytes.Buffer, size),
		max:  max,
	}
}

func (bp *BufferPool) Add(buf *bytes.Buffer) {
	if buf == nil || buf.Len() > bp.max {
		buf = newBufferL(bp.max)
	}

	bp.pool <- buf
}

func (bp *BufferPool) Cap() int {
	return cap(bp.pool)
}

func (bp *BufferPool) Clear() {
	clearing:
	for {
		select {
		case <- bp.pool:
		default:
			break clearing
		}
	}
}

func (bp *BufferPool) Fill() {
	n := bp.Cap() - bp.Len()
	for i := 0; i < n; i++ {
		bp.Add(nil)
	}
}

func (bp *BufferPool) Get() (buf *bytes.Buffer) {
	select {
	case buf = <-bp.pool:
	default:
		buf = newBufferL(bp.max)
	}

	return buf
}

func (bp *BufferPool) Len() int {
	return len(bp.pool)
}

func newBufferL(l int) *bytes.Buffer {
	s := make([]byte, l)
	return bytes.NewBuffer(s)
}
