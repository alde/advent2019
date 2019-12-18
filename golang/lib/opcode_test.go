package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpCode(t *testing.T) {
	testData := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}},
		{[]int{1101, 100, -1, 4, 0}, []int{1101, 100, -1, 4, 99}},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			input := make(chan int)
			output := make(chan int)
			halt := make(chan []int)
			defer close(input)
			defer close(output)
			defer close(halt)

			go OpCode(td.in, input, output, halt)
			assert.Equal(t, <-halt, td.out)
		})
	}
}

func Test_OpCodeWithIO(t *testing.T) {
	testData := []struct {
		in  []int
		out []int
	}{
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			input := make(chan int)
			output := make(chan int)
			halt := make(chan []int)
			defer close(input)
			defer close(output)
			defer close(halt)

			go OpCode(td.in, input, output, halt)
			input <- 1
		loop:
			for {
				select {
				case o := <-output:
					fmt.Printf("%+v\n", o)
				case res := <-halt:
					assert.Equal(t, res, td.out)
					break loop
				}
			}
		})
	}
}
