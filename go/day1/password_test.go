package main

import (
	"strings"
	"testing"
)

func TestCountZeroes(t *testing.T) {
	input := `L68
	L30
	R48
	L5
	R60
	L55
	L1
	L99
	R14
	L82`

	instructions := ReadInput(strings.NewReader(input))
	password := getPassword(instructions)

	if password != 6 {
		t.Errorf("Expected password to be 6, got %d", password)
	}
}
