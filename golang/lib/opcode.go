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
	var stepping int
	i := input[offset]
	modeSet, operation := extractOpcode(i)

	switch operation {
	case 1:
		stepping = opcodeOne(input, offset, modeSet)
	case 2:
		stepping = opcodeTwo(input, offset, modeSet)
	case 3:
		stepping = opcodeThree(input, offset, modeSet, handler)
	case 4:
		stepping = opcodeFour(input, offset, modeSet)
	default:
		return nil, fmt.Errorf("unknown operation %d", input[offset])
	}
	return record(input, handler, offset+stepping)
}

func modeSet(i int, index int, memo []int) []int {
	if index == -1 {
		return memo
	}
	memo[index] = i % 10
	return modeSet(i / 10, index - 1, memo)
}

func extractOpcode(i int) ([]int, int) {
	return modeSet(i / 100, 2, []int{0,0,0}), i % 100
}
func val(input []int, offset int, modeSet int) int {
	if modeSet == 0 {
		return input[input[offset]]
	}
	return input[offset]
}
func opcodeOne(input []int, offset int,  modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	input[c] = a + b
	return 4
}

func opcodeTwo(input []int, offset int,  modeSet []int) int {
	a := val(input, offset+1, modeSet[2])
	b := val(input, offset+2, modeSet[1])
	c := input[offset+3]
	input[c] = a * b
	return 4
}

func opcodeThree(input []int, offset int,  modeSet []int, handler func() int)  int {
	i := handler()
	target := input[offset+1]
	input[target] = i
	return 2
}

func opcodeFour(input []int, offset int,  modeSet []int)  int {
	target := input[offset+1]
	fmt.Printf("%d\n", input[target])
	return 2
}
