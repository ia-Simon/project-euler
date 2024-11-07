package main

import "testing"

func BenchmarkProblem1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem1()
	}
}
