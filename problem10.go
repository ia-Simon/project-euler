package main

import (
	"projecteuler/utils"
	"strconv"
)

func init() {
	exerciseMap[10] = Problem10
}

func Problem10() string {
	limit := 2_000_000
	sum := 2

	for i := 3; i < limit; i += 2 {
		if utils.IsPrime(i) {
			sum += i
		}
	}

	return strconv.Itoa(sum)
}
