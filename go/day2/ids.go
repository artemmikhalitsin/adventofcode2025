package main

import (
	"math"
)

func GetSumOfInvalidV1(ranges []Range) int {
	return getSumOfInvalid(ranges, isInvalidV1)
}

func GetSumOfInvalidV2(ranges []Range) int {
	return getSumOfInvalid(ranges, isInvalidV2)
}

func getSumOfInvalid(ranges []Range, isValid ValidationFunc) int {
	sum := 0
	for _, r := range ranges {
		invalidIds := findInvalid(r, isValid)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}

func findInvalid(r Range, isInvalid ValidationFunc) []int {
	var invalidIds []int
	for id := r.Min; id <= r.Max; id++ {
		if isInvalid(id) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

type ValidationFunc func(int) bool

// V1 an id is invalid if its a repeated number i.e. 11, 22, 1212, 3333, 456456
func isInvalidV1(id int) bool {
	// find order of magnitude
	order := int((math.Log10(float64(id))))

	if (order % 2) != 1 {
		return false // odd digit IDs cannot be invalid
	}
	// e.g. for 4 digit ids, mask = 100
	// so we can compare first two digits with last two digits
	mask := int(math.Pow10((order / 2) + 1))

	return (id / mask) == (id % mask)
}

// Part 2: A number is invalid if its made up of a sequence of digits
// that repeat at least twice i.e. 11111, 1212, 12121212, 123123123
func isInvalidV2(id int) bool {

	// smallest repeating sequence is length 1 (e.g. 11111)
	// largest repeating sequence is length n/2 (e.g. 1212 for 4 digit number)
	// or n/3 (e.g. 123123 for 6 digit number)
	// so we can check for all possible sequence lengths if they repeat to form the number
	numberOfDigits := getNumberOfDigits(id)
	for i := 1; i <= getMaxLengthOfRepeatingSequence(numberOfDigits); i++ {
		if numberOfDigits%i != 0 {
			continue // only consider lengths that divide evenly into numberOfDigits
		}
		mask := int(math.Pow10(i))
		sequence := id % mask

		if isRepeatedSequence(id, sequence, mask) {
			return true
		}
	}

	return false
}

func isRepeatedSequence(number int, sequence int, mask int) bool {
	// recursive base case: number == sequence
	if number == sequence {
		return true
	}
	// recursive case: check if the leading part of the number matches the sequence,
	// and then check that this is true for the rest of the number
	return number%mask == sequence && isRepeatedSequence(number/int(mask), sequence, mask)
}

func getNumberOfDigits(id int) int {
	if id == 0 {
		return 1
	}
	return int(math.Log10(float64(id))) + 1
}

func getMaxLengthOfRepeatingSequence(numDigits int) int {
	if numDigits%2 == 0 {
		return numDigits / 2 // max repeating for numbers with even length is half the length
	}
	if numDigits%3 == 0 {
		return numDigits / 3 // max repeating for numbers with length divisible by 3 is a third of the length
	}
	return 1 // number of digits is prime, so max repeating sequence is length 1
}
