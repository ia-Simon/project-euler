package main

import (
	"math"
	"strconv"
)

func init() {
	exerciseMap[3] = Problem3
}

func Problem3() string {
	num := 600_851_475_143

	if isPrime(num) {
		return strconv.Itoa(num)
	}

	for i := int(math.Sqrt(float64(num))); i > 1; i-- {
		if num%i == 0 && isPrime(i) {
			return strconv.Itoa(i)
		}
	}
	return strconv.Itoa(num)
}

func isPrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
