package day6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAllOrbits(t *testing.T) {
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	planets := parse(input)
	actual := planets.countAllOrbits()
	assert.Equal(t, actual, 42)
}

func Test_getPath(t *testing.T) {
	testData := []struct {
		in  string
		out []string
	}{
		{"C", []string{"B", "COM"}},
		{"YOU", []string{"K", "J", "E", "D", "C", "B", "COM"}},
	}
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	planets := parse(input)
	for _, td := range testData {
		t.Run(td.in, func(t *testing.T) {
			path := planets.getPath(td.in)
			assert.Equal(t, path, td.out)
		})
	}
}

func Test_getDistance(t *testing.T) {
	testData := []struct {
		from     string
		to       string
		distance int
	}{
		{"L", "H", 6},
		{"YOU", "SAN", 4},
	}
	input := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	planets := parse(input)
	for _, td := range testData {
		t.Run(fmt.Sprintf("distance from %s to %s", td.from, td.to), func(t *testing.T) {
			distance := planets.getDistance(td.from, td.to)
			assert.Equal(t, distance, td.distance)
		})
	}
}
