package main

import (
	"strconv"
)

func init() {
	exerciseMap[2] = Problem2
}

func Problem2() string {
	minus2 := 0
	minus1 := 1
	accum := 0
	for {
		minus2, minus1 = minus1, minus2+minus1
		if minus1 > 4_000_000 {
			break
		}

		if minus1%2 == 0 {
			accum += minus1
		}
	}

	return strconv.Itoa(accum)
}
