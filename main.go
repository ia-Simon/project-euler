package main

import (
	"flag"
	"fmt"
	"log"
)

type ProblemFunc func() string

var (
	fProblemNumber = 0

	exerciseMap map[int]ProblemFunc = make(map[int]ProblemFunc)
)

func init() {
	flag.IntVar(&fProblemNumber, "n", fProblemNumber, "Specifies which problem to run.")
}

func main() {
	flag.Parse()

	if fProblemNumber <= 0 {
		log.Fatal("invalid exercise number; must be greater than 0")
	}

	exFunc, ok := exerciseMap[fProblemNumber]
	if !ok {
		log.Fatalf("exercise #%d not found\n", fProblemNumber)
	}

	fmt.Println("Result: ", exFunc())
}
