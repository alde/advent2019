package main

import (
	"fmt"
	"testing"
)

func Test_partOneValidator(t *testing.T) {
	var testData = []struct {
		in  int
		out bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			actual := partOneValidator(td.in)
			if actual != td.out {
				t.Fail()
			}
		})
	}
}

func Test_partTwoValidator(t *testing.T) {
	var testData = []struct {
		in  int
		out bool
	}{
		{111111, false},
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("%+v", td.in), func(t *testing.T) {
			actual := partTwoValidator(td.in)
			if actual != td.out {
				t.Fail()
			}
		})
	}
}
