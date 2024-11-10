package utils

import "math"

func IsPrime(n int) bool {
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
