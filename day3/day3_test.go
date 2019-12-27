package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Manhattan(t *testing.T) {
	testData := []struct {
		in  [][]string
		out int
	}{
		{
			[][]string{
				[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
				[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			}, 159,
		},
		{
			[][]string{
				[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
				[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			}, 135,
		},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual, _ := manhattan(td.in)
			assert.Equal(t, td.out, actual)
		})
	}
}

func Test_FewestCombinedSteps(t *testing.T) {
	testData := []struct {
		in  [][]string
		out int
	}{
		{
			[][]string{
				[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
				[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			}, 610,
		},
		{
			[][]string{
				[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
				[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			}, 410,
		},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			_, actual := manhattan(td.in)
			assert.Equal(t, td.out, actual)
		})
	}
}
