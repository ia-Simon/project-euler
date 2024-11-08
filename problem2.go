package main

import "strconv"

func init() {
	exerciseMap[2] = Problem2
}

func Problem2() string {

	return strconv.Itoa(fib(40))
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
