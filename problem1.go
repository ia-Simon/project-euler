package main

import "strconv"

func init() {
	exerciseMap[1] = Problem1
}

func Problem1() string {
	accum := 0
	for i := range 10 {
		switch {
		case i%3 == 0:
			accum += i
		case i%5 == 0:
			accum += i
		}
	}

	return strconv.Itoa(accum)
}
