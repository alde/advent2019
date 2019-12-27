package day2

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_GravityAssist(t *testing.T) {
	testData := []struct {
		in  []int
		out int
	}{
		{[]int{1,0,0,0,99}, 2},
		{[]int{2,3,0,3,99}, 2},
		{[]int{2,4,4,5,99,0}, 2},
		{[]int{1,1,1,4,99,5,6,0,99}, 30},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := gravityAssist(td.in)
			assert.Equal(t, td.out, actual)
		})
	}
}
