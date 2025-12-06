package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("The solution is", solve(input))
}

func solve(input []byte) int {
	ranges := extractRanges(string(input))
	return getSumOfInvalid(ranges)
}
