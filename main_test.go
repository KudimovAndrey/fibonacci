package main

import (
	"testing"
)

func Benchmark_firstFibonacci(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		f := firstFibonacci()
		for i := 1; i <= 10; i++ {
			f()
		}
	}
	b.StopTimer()
}

func Benchmark_secondFibonacci(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i <= 10; i++ {
			secondFibonacci(i)
		}
	}
	b.StopTimer()
}
