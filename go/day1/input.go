package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Direction int

const (
	L Direction = iota
	R
)

type Instruction struct {
	Direction Direction
	Clicks    int
}

func ReadInput(input io.Reader) []Instruction {
	sc := bufio.NewScanner(input)
	var result []Instruction
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		result = append(result, ParseInstruction(line))
	}
	return result
}

func makeDirection(dir byte) Direction {
	if dir == 'L' {
		return L
	}
	if dir == 'R' {
		return R
	}
	panic(fmt.Sprintf("unknown direction: %c", dir))
}

func ParseInstruction(line string) Instruction {
	dir := makeDirection(line[0])
	clicks, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(fmt.Sprintf("cannot parse clicks from line: %s", line))
	}

	return Instruction{Direction: dir, Clicks: clicks}
}
