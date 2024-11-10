package main

import "strconv"

func init() {
	exerciseMap[6] = Problem6
}

func Problem6() string {
	sqSum := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		sqSum += (i * i)
		sum += i
	}

	return strconv.Itoa((sum * sum) - sqSum)
}
