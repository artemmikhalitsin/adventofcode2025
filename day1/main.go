package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	input := ReadInput(inputFile)
	fmt.Println("The solution is", getPassword(input))
}
