// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes_test

import (
	"bytes"
	"testing"
	"time"

	ntbytes "go.ntrrg.dev/ntgo/bytes"
)

// nolint
const (
	BUF_MAX_COUNT = 20 // Must be higher than 5
	BUF_MAX_SIZE  = 10 // Must be higher than 1
)

func TestNewBufferPool(t *testing.T) {
	t.Parallel()

	n := BUF_MAX_COUNT
	bp := ntbytes.NewBufferPool(n, BUF_MAX_SIZE)

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty")
	}

	if bp.Cap() != n {
		t.Errorf("the buffer pool capacity should be %d, got %d", n, bp.Cap())
	}
}

func TestBufferPool_Add(t *testing.T) {
	t.Parallel()

	n := BUF_MAX_COUNT
	max := BUF_MAX_SIZE
	bp := ntbytes.NewBufferPool(n, max)
	bp.Fill()
	bp.Get() // Discard first buffer

	// Reuse buffer if there is room

	bp.Add(bytes.NewBuffer(make([]byte, max-5)))

	for i := 1; i < n; i++ { // Dicard all buffers but last
		bp.Get()
	}

	buf := bp.Get()

	if buf.Cap() != max-5 {
		t.Errorf("the buffer pool didn't reuse the buffer")
	}

	// Overflow pool

	bp.Fill()
	bp.Add(bytes.NewBuffer(make([]byte, max-5)))

	for i := 1; i < n; i++ { // Dicard all buffers but last
		bp.Get()
	}

	buf = bp.Get()

	if buf.Cap() == max-5 {
		t.Errorf("the buffer pool reused the buffer even when there was no room")
	}
}

func BenchmarkBufferPool_Add(b *testing.B) {
	max := BUF_MAX_SIZE
	bp := ntbytes.NewBufferPool(1, max)

	cases := []struct {
		name      string
		newBuffer func() *bytes.Buffer
	}{
		{"Nil", func() *bytes.Buffer { return nil }},

		{"Valid", func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, max))
		}},

		{"Invalid", func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, max+5))
		}},
	}

	for _, c := range cases {
		c := c

		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				bp.Clear()
				buf := c.newBuffer()
				b.StartTimer()

				bp.Add(buf)
			}
		})
	}
}

func TestBufferPool_AddWait(t *testing.T) {
	t.Parallel()

	max := BUF_MAX_SIZE
	bp := ntbytes.NewBufferPool(BUF_MAX_COUNT, max)

	// Let the pool create a buffer with bp.max bytes

	bp.AddWait(nil)
	buf := bp.Get()

	if buf.Cap() != max {
		t.Errorf("the buffer pool created a buffer with bad size")
	}

	// Discard buffer with more than bp.max bytes

	bp.AddWait(bytes.NewBuffer(make([]byte, max+5)))
	buf = bp.Get()

	if buf.Cap() != max {
		t.Errorf("the buffer pool reused a buffer with bad size")
	}

	// Reuse valid buffer

	bp.AddWait(bytes.NewBuffer(make([]byte, 0, max-5)))
	buf = bp.Get()

	if buf.Cap() != max-5 {
		t.Errorf("the buffer pool didn't reuse the buffer")
	}

	// Overflow pool

	bp.Fill()

	done := make(chan struct{})

	go func(done chan<- struct{}) {
		bp.AddWait(nil)
		done <- struct{}{}
	}(done)

	select {
	case <-done:
		t.Errorf("the buffer pool didn't wait for appending the buffer")
	default:
	}
}

func TestBufferPool_Clear(t *testing.T) {
	t.Parallel()

	bp := ntbytes.NewBufferPool(BUF_MAX_COUNT, BUF_MAX_SIZE)
	bp.Fill()
	bp.Clear()

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty after calling Clear")
	}
}

func TestBufferPool_Fill(t *testing.T) {
	t.Parallel()

	bp := ntbytes.NewBufferPool(BUF_MAX_COUNT, BUF_MAX_SIZE)
	bp.Fill()

	if bp.Len() < bp.Cap() {
		t.Errorf("the buffer pool should be full after calling Clear")
	}
}

func TestBufferPool_Get(t *testing.T) {
	t.Parallel()

	max := BUF_MAX_SIZE
	bp := ntbytes.NewBufferPool(BUF_MAX_COUNT, max)

	// Let the pool create a buffer with bp.max bytes

	if bp.Len() > 0 {
		t.Errorf("the buffer pool should be empty")
	}

	buf := bp.Get()

	if buf.Cap() != max {
		t.Errorf("the buffer pool created a buffer with bad size")
	}

	// Reuse valid buffer

	bp.Add(bytes.NewBuffer(make([]byte, 0, max-5)))
	buf = bp.Get()

	if buf.Cap() != max-5 {
		t.Errorf("the buffer pool didn't reuse the buffer")
	}
}

func TestBufferPool_GetWait(t *testing.T) {
	t.Parallel()

	bp := ntbytes.NewBufferPool(BUF_MAX_COUNT, BUF_MAX_SIZE)
	bp.Add(nil)

	// Use available buffer

	c := make(chan *bytes.Buffer)

	go func(c chan<- *bytes.Buffer) {
		c <- bp.GetWait()
	}(c)

	select {
	case <-c:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("the buffer pool didn't use available buffers")
	}

	// Wait for a buffer

	c = make(chan *bytes.Buffer)

	go func(c chan<- *bytes.Buffer) {
		c <- bp.GetWait()
	}(c)

	select {
	case <-c:
		t.Errorf("the buffer pool don't wait for available buffers")
	case <-time.After(10 * time.Millisecond):
	}
}
