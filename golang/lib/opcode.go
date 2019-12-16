package lib
import (
	"fmt"
)
// ResetCode resets the input code
func ResetCode(input []int, noun int, verb int) ([]int, error) {
	input[0] = noun
	input[1] = verb
	return OpCode(input)
}

// OpCode runs the computer
func OpCode(input []int) ([]int, error) {
	return record(input, 0)
}

func record(input []int, offset int) ([]int, error) {
	if input[offset] == 99 {
		return input, nil
	}
	a := input[offset+1]
	b := input[offset+2]
	c := input[offset+3]
	switch input[offset] {
	case 1:
		input[c] = input[a] + input[b]
	case 2:
		input[c] = input[a] * input[b]
	default:
		return nil, fmt.Errorf("unknown operation %d", input[offset])
	}
	return record(input, offset+4)
}
