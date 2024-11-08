package main

import "testing"

func BenchmarkProblem2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem2()
	}
}
