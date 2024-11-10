package main

import (
	"projecteuler/utils"
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
			if utils.IsPalindrome(i*j) && i*j > highest {
				highest = i * j
			}
		}
	}

	return strconv.Itoa(highest)
}
