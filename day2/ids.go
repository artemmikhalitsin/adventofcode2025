package main

import "math"

func getSumOfInvalid(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		invalidIds := findInvalid(r)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}

func findAllInvalid(ranges []Range) []int {
	var allInvalid []int
	for _, r := range ranges {
		invalidIds := findInvalid(r)
		allInvalid = append(allInvalid, invalidIds...)
	}
	return allInvalid
}

func findInvalid(r Range) []int {
	var invalidIds []int
	for id := r.Min; id <= r.Max; id++ {
		// an id is invalid if its a repeated number
		// find order of magnitude
		order := int((math.Log10(float64(id))))

		if (order % 2) != 1 {
			continue // odd digit IDs cannot be invalid
		}
		// e.g. for 4 digit ids, mask = 100
		// so we can compare first two digits with last two digits
		mask := int(math.Pow10((order / 2) + 1))

		if (id / mask) == (id % mask) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}
