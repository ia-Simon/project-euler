package main

import (
	"math"
	"projecteuler/utils"
	"strconv"
)

func init() {
	exerciseMap[3] = Problem3
}

func Problem3() string {
	num := 600_851_475_143

	if utils.IsPrime(num) {
		return strconv.Itoa(num)
	}

	for i := int(math.Sqrt(float64(num))); i > 1; i-- {
		if num%i == 0 && utils.IsPrime(i) {
			return strconv.Itoa(i)
		}
	}
	return strconv.Itoa(num)
}
