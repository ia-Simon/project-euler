package utils

import "strconv"

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
