package utils

func IsPrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
