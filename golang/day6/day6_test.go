package main

import (
	"testing"
	  "github.com/stretchr/testify/assert"
)

func Test_countAllOrbits(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	planets := parse(input)
	actual := planets.countAllOrbits()
	assert.Equal(t, 42, actual)
}