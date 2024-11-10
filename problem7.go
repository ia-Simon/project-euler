package main

import (
	"projecteuler/utils"
	"strconv"
)

func init() {
	exerciseMap[7] = Problem7
}

func Problem7() string {
	count := 0
	for i := 2; ; i++ {
		if utils.IsPrime(i) {
			count++
		}

		if count == 10_001 {
			return strconv.Itoa(i)
		}
	}
}
