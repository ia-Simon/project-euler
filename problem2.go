package main

import (
	"strconv"
)

func init() {
	exerciseMap[2] = Problem2
}

func Problem2() string {

	// return strconv.Itoa(fib(42))
	// return strconv.Itoa(tailFib(42))
	return strconv.Itoa(memoize(fib)(42))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func tailFib(n int) int {
	last := 0
	accum := 1
	for range n - 1 {
		last, accum = accum, last+accum
	}

	return accum
}

func memoize[In comparable, Out any](fn func(In) Out) func(In) Out {
	cache := make(map[In]Out)

	return func(in In) Out {
		if out, ok := cache[in]; ok {
			return out
		}

		out := fn(in)
		cache[in] = out
		return out
	}
}
