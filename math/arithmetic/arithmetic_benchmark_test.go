// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package arithmetic_test

import (
	"fmt"
	"testing"

	a "github.com/ntrrg/ntgo/math/arithmetic"
)

// See common_test.go

func benchmarkAdd(n int, b *testing.B) {
	operanders := make([]interface{}, n)

	for i := range operanders {
		operanders[i] = i
	}

	for i := 0; i <= b.N; i++ {
		a.Add(operanders...)
	}
}

func BenchmarkAdd_2(b *testing.B)    { benchmarkAdd(2, b) }
func BenchmarkAdd_20(b *testing.B)   { benchmarkAdd(20, b) }
func BenchmarkAdd_200(b *testing.B)  { benchmarkAdd(200, b) }
func BenchmarkAdd_2000(b *testing.B) { benchmarkAdd(2000, b) }

func benchmarkDiv(n int, b *testing.B) {
	o := make([]Operand, n)

	for i := 0; i < n; i++ {
		o[i] = Operand(fmt.Sprintf("%b", i))
	}

	operanders := Operanders(o)

	for i := 0; i <= b.N; i++ {
		a.Div(operanders...)
	}
}

func BenchmarkDiv_2(b *testing.B)    { benchmarkDiv(2, b) }
func BenchmarkDiv_20(b *testing.B)   { benchmarkDiv(20, b) }
func BenchmarkDiv_200(b *testing.B)  { benchmarkDiv(200, b) }
func BenchmarkDiv_2000(b *testing.B) { benchmarkDiv(2000, b) }

func benchmarkMul(n int, b *testing.B) {
	o := make([]Operand, n)

	for i := 0; i < n; i++ {
		o[i] = Operand(fmt.Sprintf("%b", i))
	}

	operanders := Operanders(o)

	for i := 0; i <= b.N; i++ {
		a.Mul(operanders...)
	}
}

func BenchmarkMul_2(b *testing.B)    { benchmarkMul(2, b) }
func BenchmarkMul_20(b *testing.B)   { benchmarkMul(20, b) }
func BenchmarkMul_200(b *testing.B)  { benchmarkMul(200, b) }
func BenchmarkMul_2000(b *testing.B) { benchmarkMul(2000, b) }

func benchmarkSub(n int, b *testing.B) {
	o := make([]Operand, n)

	for i := 0; i < n; i++ {
		o[i] = Operand(fmt.Sprintf("%b", i))
	}

	operanders := Operanders(o)

	for i := 0; i <= b.N; i++ {
		a.Sub(operanders...)
	}
}

func BenchmarkSub_2(b *testing.B)    { benchmarkSub(2, b) }
func BenchmarkSub_20(b *testing.B)   { benchmarkSub(20, b) }
func BenchmarkSub_200(b *testing.B)  { benchmarkSub(200, b) }
func BenchmarkSub_2000(b *testing.B) { benchmarkSub(2000, b) }
