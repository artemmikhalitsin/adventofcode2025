package main

func getPassword(instructions []Instruction) int {
	position := 50
	password := 0
	for _, instr := range instructions {
		password += countZeroes(position, instr.Direction, instr.Clicks)
		position = rotate(position, instr.Direction, instr.Clicks)
	}

	return password
}

func countZeroes(initial int, direction Direction, clicks int) int {
	zeroes := clicks / 100 // each 100 clicks adds one zero
	remainingClicks := clicks % 100
	if rotationPassesZero(initial, direction, remainingClicks) {
		zeroes++
	}
	return zeroes
}

func rotationPassesZero(position int, direction Direction, clicks int) bool {
	if position == 0 {
		return false // if initial position is 0 then only full rotations return it to zero
	}
	if direction == R {
		return clicks >= (100 - position) // going right, must click at least (100 - position) to reach zero
	}
	return clicks >= position // going left, must click at least position to reach zero
}

func rotate(position int, direction Direction, clicks int) int {
	if direction == R {
		return (position + clicks) % 100
	}
	// direction == L
	newPosition := position - clicks
	for newPosition < 0 {
		newPosition += 100
	}
	return newPosition
}
