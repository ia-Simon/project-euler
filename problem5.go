package main

import "strconv"

func init() {
	exerciseMap[5] = Problem5
}

func Problem5() string {
	for i := 1; ; i++ {
		if divisibleBy(i, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}) {
			return strconv.Itoa(i)
		}
	}
}

func divisibleBy(num int, divisors []int) bool {
	for _, v := range divisors {
		if num%v != 0 {
			return false
		}
	}

	return true
}
