package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	testData := []struct {
		in  string
		out [][]rune
	}{
		{
			".#..#\n.....\n#####\n....#\n...##",
			[][]rune{
				[]rune{'.', '#', '.', '.', '#'},
				[]rune{'.', '.', '.', '.', '.'},
				[]rune{'#', '#', '#', '#', '#'},
				[]rune{'.', '.', '.', '.', '#'},
				[]rune{'.', '.', '.', '#', '#'},
			},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := parse(td.in)
			assert.Equal(t, actual, td.out)
		})
	}
}

func Test_countVisible(t *testing.T) {
	testData := []struct {
		astr []coordinate
		out  int
		pos  coordinate
	}{
		{
			[]coordinate{
				{1, 0}, {4, 0}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {4, 3}, {3, 4}, {4, 4},
			},
			8,
			coordinate{3, 4},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			aC, aPos := countVisible(td.astr)
			assert.Equal(t, td.out, aC)
			assert.Equal(t, td.pos, aPos)
		})
	}
}
