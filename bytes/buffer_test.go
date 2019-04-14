// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes_test

import (
	"bytes"
	"testing"

	ntbytes "github.com/ntrrg/ntgo/bytes"
)

func TestNewBufferPool(t *testing.T) {
	c := 50
	bp := ntbytes.NewBufferPool(c, 10)

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty")
	}

	if bp.Cap() != c {
		t.Errorf("the buffer pool capacity should be %d, got %d", c, bp.Cap())
	}
}

func TestBufferPool_Add(t *testing.T) {
	max := 10
	bp := ntbytes.NewBufferPool(50, max)

	// Let the pool create a buffer with bp.max bytes

	bp.Add(nil)
	buf := bp.Get()

	if buf.Len() != max {
		t.Errorf("the buffer pool created a buffer with bad size")
	}

	// Discard buffer with more than bp.max bytes

	bp.Add(bytes.NewBuffer(make([]byte, 15)))
	buf = bp.Get()

	if buf.Len() != max {
		t.Errorf("the buffer pool reused a buffer with bad size")
	}

	// Reuse valid buffer

	bp.Add(bytes.NewBuffer(make([]byte, 5)))
	buf = bp.Get()

	if buf.Len() != 5 {
		t.Errorf("the buffer pool didn't reused the buffer")
	}
}

func TestBufferPool_Clear(t *testing.T) {
	bp := ntbytes.NewBufferPool(50, 10)
	bp.Fill()
	bp.Clear()

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty after calling Clear")
	}
}

func TestBufferPool_Fill(t *testing.T) {
	bp := ntbytes.NewBufferPool(50, 10)
	bp.Fill()

	if bp.Len() < bp.Cap() {
		t.Errorf("the buffer pool should be full after calling Clear")
	}
}

func TestBufferPool_Get(t *testing.T) {
	bp := ntbytes.NewBufferPool(50, 10)

	// New buffer

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty")
	}

	buf := bp.Get()

	// Reused buffer

	bp.Add(buf)

	if bp.Len() < 1 {
		t.Errorf("the buffer pool should have at least one buffer")
	}

	_ = bp.Get()
}
