// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes_test

import (
	"testing"
	"time"

	ntbytes "go.ntrrg.dev/ntgo/bytes"
)

// nolint
const (
	SLICE_MAX_COUNT = 20 // Must be higher than 5
	SLICE_MAX_SIZE  = 10 // Must be higher than 1
)

func TestNewPool(t *testing.T) {
	t.Parallel()

	n := SLICE_MAX_COUNT
	p := ntbytes.NewPool(n, SLICE_MAX_SIZE)

	if p.Len() > 0 {
		t.Errorf("the slice pool should be empty")
	}

	if p.Cap() != n {
		t.Errorf("the slice pool capacity should be %d, got %d", n, p.Cap())
	}
}

func TestPool_Add(t *testing.T) {
	t.Parallel()

	n := SLICE_MAX_COUNT
	max := SLICE_MAX_SIZE
	p := ntbytes.NewPool(n, max)
	p.Fill()
	p.Get() // Discard first slice

	// Reuse slice if there is room

	p.Add(make([]byte, max-5))

	for i := 1; i < n; i++ { // Dicard all slices but last
		p.Get()
	}

	s := p.Get()

	if cap(s) != max-5 {
		t.Errorf("the slice pool didn't reuse the slice")
	}

	// Overflow pool

	p.Fill()
	p.Add(make([]byte, max-5))

	for i := 1; i < n; i++ { // Dicard all slices but last
		p.Get()
	}

	s = p.Get()

	if cap(s) == max-5 {
		t.Errorf("the slice pool reused the slice even when there was no room")
	}
}

func BenchmarkPool_Add(b *testing.B) {
	max := SLICE_MAX_SIZE
	p := ntbytes.NewPool(1, max)

	cases := []struct {
		name     string
		newSlice func() []byte
	}{
		{"Nil", func() []byte { return nil }},

		{"Valid", func() []byte {
			return make([]byte, max)
		}},

		{"Invalid", func() []byte {
			return make([]byte, max+5)
		}},
	}

	for _, c := range cases {
		c := c

		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				p.Clear()
				s := c.newSlice()
				b.StartTimer()

				p.Add(s)
			}
		})
	}
}

func TestPool_AddWait(t *testing.T) {
	t.Parallel()

	max := SLICE_MAX_SIZE
	p := ntbytes.NewPool(SLICE_MAX_COUNT, max)

	// Let the pool create a slice with p.max bytes

	p.AddWait(nil)
	s := p.Get()

	if cap(s) != max {
		t.Errorf("the slice pool created a slice with bad size")
	}

	// Discard slice with more than p.max bytes

	p.AddWait(make([]byte, max+5))
	s = p.Get()

	if cap(s) != max {
		t.Errorf("the slice pool reused a slice with bad size")
	}

	// Reuse valid slice

	p.AddWait(make([]byte, max-5))
	s = p.Get()

	if cap(s) != max-5 {
		t.Errorf("the slice pool didn't reuse the slice")
	}

	// Overflow pool

	p.Fill()

	done := make(chan struct{})

	go func(done chan<- struct{}) {
		p.AddWait(nil)
		done <- struct{}{}
	}(done)

	select {
	case <-done:
		t.Errorf("the slice pool didn't wait for appending the slice")
	default:
	}
}

func TestPool_Clear(t *testing.T) {
	t.Parallel()

	p := ntbytes.NewPool(SLICE_MAX_COUNT, SLICE_MAX_SIZE)
	p.Fill()
	p.Clear()

	if p.Len() > 0 {
		t.Errorf("the slice pool should be empty after calling Clear")
	}
}

func TestPool_Fill(t *testing.T) {
	t.Parallel()

	p := ntbytes.NewPool(SLICE_MAX_COUNT, SLICE_MAX_SIZE)
	p.Fill()

	if p.Len() < p.Cap() {
		t.Errorf("the slice pool should be full after calling Clear")
	}
}

func TestPool_Get(t *testing.T) {
	t.Parallel()

	max := SLICE_MAX_SIZE
	p := ntbytes.NewPool(SLICE_MAX_COUNT, max)

	// Let the pool create a slice with p.max bytes

	if p.Len() > 0 {
		t.Errorf("the slice pool should be empty")
	}

	s := p.Get()

	if cap(s) != max {
		t.Errorf("the slice pool created a slice with bad size")
	}

	// Reuse valid slice

	p.Add(make([]byte, max-5))
	s = p.Get()

	if cap(s) != max-5 {
		t.Errorf("the slice pool didn't reuse the slice")
	}
}

func TestPool_GetWait(t *testing.T) {
	t.Parallel()

	p := ntbytes.NewPool(SLICE_MAX_COUNT, SLICE_MAX_SIZE)
	p.Add(nil)

	// Use available slice

	c := make(chan []byte)

	go func(c chan<- []byte) {
		c <- p.GetWait()
	}(c)

	select {
	case <-c:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("the slice pool didn't use available slices")
	}

	// Wait for a slice

	c = make(chan []byte)

	go func(c chan<- []byte) {
		c <- p.GetWait()
	}(c)

	select {
	case <-c:
		t.Errorf("the slice pool don't wait for available slices")
	case <-time.After(10 * time.Millisecond):
	}
}
