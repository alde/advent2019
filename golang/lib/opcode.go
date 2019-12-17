package lib

import (
	"fmt"
)

// ResetCode resets the input code
func ResetCode(input []int, noun int, verb int) ([]int, error) {
	input[0] = noun
	input[1] = verb

	return OpCode(input, func() (i int) {
		fmt.Scan(&i)
		return i
	})
}

// OpCode runs the computer
func OpCode(input []int, handler func() int) ([]int, error) {
	return record(input, handler, 0)
}

func record(input []int, handler func() int, offset int) ([]int, error) {
	if input[offset] == 99 {
		return input, nil
	}
	var pointer int
	i := input[offset]
	modeSet, operation := extractOpcode(i)

	switch operation {
	case 1:
		pointer = addition(input, offset, modeSet)
	case 2:
		pointer = multiplication(input, offset, modeSet)
	case 3:
		pointer = readInput(input, offset, modeSet, handler)
	case 4:
		pointer = output(input, offset, modeSet)
	case 5:
		pointer = jumpIfTrue(input, offset, modeSet)
	case 6:
		pointer = jumpIfFalse(input, offset, modeSet)
	case 7:
		pointer = lessThan(input, offset, modeSet)
	case 8:
		pointer = equals(input, offset, modeSet)
	default:
		return nil, fmt.Errorf("unknown operation %d", input[offset])
	}
	return record(input, handler, pointer)
}

func modeSet(i int, index int, memo []int) []int {
	if index == -1 {
		return memo
	}
	memo[index] = i % 10
	return modeSet(i/10, index-1, memo)
}

func extractOpcode(i int) ([]int, int) {
	return modeSet(i/100, 2, []int{0, 0, 0}), i % 100
}
func val(input []int, offset int, modeSet int) int {
	if modeSet == 0 {
		return input[input[offset]]
	}
	return input[offset]
}
func addition(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	input[c] = a + b
	return offset + 4
}

func multiplication(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	input[c] = a * b
	return offset + 4
}

func readInput(input []int, offset int, modeSet []int, handler func() int) int {
	i := handler()
	target := input[offset+1]
	input[target] = i
	return offset + 2
}

func output(input []int, offset int, modeSet []int) int {
	target := input[offset+1]
	fmt.Printf("%d\n", input[target])
	return offset + 2
}

func jumpIfTrue(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	if a != 0 {
		return b
	}
	return offset + 3
}

func jumpIfFalse(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	if a == 0 {
		return b
	}
	return offset + 3
}

func lessThan(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	if a < b {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return offset + 4
}

func equals(input []int, offset int, modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	if a == b {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return offset + 4
}
