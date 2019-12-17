package lib

import (
	"fmt"
	"testing"
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
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			actual, _ := OpCode(td.in, func() int {
				return 1
			}, func(i int) {
				fmt.Printf("%d\n", i)
			})
			validate(t, actual, td.out)
		})
	}
}

func Test_extractOpcode(t *testing.T) {
	testData := []struct {
		in     int
		mode   []int
		opcode int
	}{
		{2, []int{0, 0, 0}, 2},
		{1002, []int{0, 1, 0}, 2},
	}
	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			mode, code := extractOpcode(td.in)
			validate(t, mode, td.mode)
			if code != td.opcode {
				fmt.Printf("expected: %+v\nactual: %+v\n", td.opcode, code)
				t.Fail()
			}
		})
	}
}

func validate(t *testing.T, actual []int, expected []int) {
	t.Logf("validating %+v == %+v\n", expected, actual)
	for idx, e := range actual {
		if expected[idx] != e {
			fmt.Printf("expected: %+v\nactual: %+v\n", expected, actual)
			t.Fail()
		}
	}
}
