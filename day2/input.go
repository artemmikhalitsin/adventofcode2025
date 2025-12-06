package main

import (
	"fmt"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func extractRanges(input string) []Range {
	var ranges []Range
	chunks := strings.Split(input, ",")

	for _, chunk := range chunks {
		ranges = append(ranges, ReadRange(chunk))
	}
	return ranges
}

func ReadRange(r string) Range {
	var min, max int
	n, err := fmt.Sscanf(r, "%d-%d", &min, &max)
	if err != nil {
		panic("cannot parse range: " + err.Error())
	}
	if n != 2 {
		panic("cannot parse range: " + r)
	}
	return Range{Min: min, Max: max}
}
