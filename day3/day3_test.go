package day3

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_Manhattan(t *testing.T) {
	testData := []struct {
		in  [][]string
		out int
	}{
		{
			[][]string{
				[]string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"},
				[]string{"U62","R66","U55","R34","D71","R55","D58","R83"},
			}, 159,
		},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := manhattan(td.in)
			assert.Equal(t, td.out, actual)
		})
	}
}
