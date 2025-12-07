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

	fmt.Println("The part one solution is", solvePartOne(input))
	fmt.Println("The part two solution is", solvePartTwo(input))
}

func solvePartOne(input []byte) int {
	ranges := extractRanges(string(input))
	return GetSumOfInvalidV1(ranges)
}

func solvePartTwo(input []byte) int {
	ranges := extractRanges(string(input))
	return GetSumOfInvalidV2(ranges)
}
