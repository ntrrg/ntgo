// Copyright 2021 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package bytes

// Pool is a pool implementation of byte slices. This allows to reuse slices,
// which means less allocations. Slices are guaranteed to live in memory, which
// is a key difference between sync.Pool.
type Pool struct {
	pool chan []byte
	max  int
}

// NewPool creates a pool that will hold up to n slices. Any new slice will be
// created with up to max bytes.
func NewPool(n, max int) Pool {
	return Pool{
		pool: make(chan []byte, n),
		max:  max,
	}
}

// Add appends s to the pool. Add behaves like AddWait, but if there is no more
// room for slices, s will be discarded.
func (p *Pool) Add(s []byte) {
	if p.Len() == p.Cap() {
		return
	}

	p.AddWait(s)
}

// AddWait appends s to the pool. If there is no room for more slices, AddWait
// will wait until s can be appended. If s is nil or has more than p.max
// bytes, it will be discarded and a new slice will be allocated. After calling
// AddWait, s must not be used anymore.
func (p *Pool) AddWait(s []byte) {
	if s == nil || cap(s) > p.max {
		s = make([]byte, 0, p.max)
	} else {
		s = s[:0]
	}

	p.pool <- s
}

// Cap returns the pool capacity.
func (p *Pool) Cap() int {
	return cap(p.pool)
}

// Clear discards all the slices in the pool.
func (p *Pool) Clear() {
	for len(p.pool) > 0 {
		<-p.pool
	}
}

// Fill fills the pool up to its available capacity.
func (p *Pool) Fill() {
	n := p.Cap() - p.Len()

	for i := 0; i < n; i++ {
		p.Add(nil)
	}
}

// Get returns a slice from the pool. If there are no slices in the pool, a new
// one with a capacity of p.max bytes will be returned.
func (p *Pool) Get() (s []byte) {
	select {
	case s = <-p.pool:
	default:
		s = make([]byte, 0, p.max)
	}

	return s
}

// GetWait returns a slice from the pool. If there are no slices in the pool,
// this will wait until one is available.
func (p *Pool) GetWait() []byte {
	return <-p.pool
}

// Len returns the amount of slices in the pool.
func (p *Pool) Len() int {
	return len(p.pool)
}
