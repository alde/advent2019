package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLayers(t *testing.T) {
	testData := []struct {
		in     string
		width  int
		height int
		out    [][]int
	}{
		{
			"123456789012", 3, 2, [][]int{[]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 0, 1, 2}},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := parseLayers(td.in, td.width, td.height)
			assert.Equal(t, actual, td.out)
		})
	}
}

func Test_checksum(t *testing.T) {
	testData := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{[]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 0, 1, 2}}, 1,
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := checksum(td.in)
			assert.Equal(t, actual, td.out)
		})
	}
}

func Test_flatten(t *testing.T) {
	testData := []struct {
		in  string
		out []int
	}{
		{
			"0222112222120000", []int{0, 1, 1, 0},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			parsed := parseLayers(td.in, 2, 2)
			actual := flatten(parsed, 2+2)
			assert.Equal(t, td.out, actual)
		})
	}
}
