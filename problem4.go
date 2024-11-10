package main

import (
	"strconv"
)

func init() {
	exerciseMap[4] = Problem4
}

func Problem4() string {
	num := 999

	highest := 0
	for j := num; j >= 100; j-- {
		for i := num; i >= j; i-- {
			if isPalindrome(i*j) && i*j > highest {
				highest = i * j
			}
		}
	}

	return strconv.Itoa(highest)
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
