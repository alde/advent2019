package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fuel(t *testing.T) {
	testData := []struct {
		in  []int
		out int
	}{
		{[]int{12}, 2},
		{[]int{14}, 2},
		{[]int{1969}, 654},
		{[]int{100756}, 33583},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := totalFuel(td.in, 0)
			assert.Equal(t, td.out, actual)
		})
	}
}

func Test_FuelAccountingForFuel(t *testing.T) {
	testData := []struct {
		in  []int
		out int
	}{
		{[]int{12}, 2},
		{[]int{14}, 2},
		{[]int{1969}, 966},
		{[]int{100756}, 50346},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			actual := totalFuelInc(td.in, 0)
			assert.Equal(t, td.out, actual)
		})
	}
}
