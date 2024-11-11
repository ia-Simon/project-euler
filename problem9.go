package main

import (
	"math"
	"strconv"
)

func init() {
	exerciseMap[9] = Problem9
}

func Problem9() string {
	max := 1000

	for a := 1; a <= max; a++ {
		for b := max; b > a; b-- {
			sqrt := math.Sqrt(float64(a*a + b*b))

			if sqrt != math.Trunc(sqrt) {
				continue
			}

			c := int(sqrt)
			if c <= b {
				break
			}

			if c+b+a == 1000 {
				return strconv.Itoa(c * b * a)
			}
		}
	}

	return "Not found"
}
