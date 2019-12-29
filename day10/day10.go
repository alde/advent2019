package day10

import (
	"math"
	"strings"

	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := lib.ReadFile("day10/input.txt")
	amap := parse(input)
	asteroids := asteroids(amap)

	count, pos := countVisible(asteroids)

	logrus.WithFields(logrus.Fields{
		"x":       pos.x,
		"y":       pos.y,
		"visible": count,
	}).Info("best position for base")
}

func parse(s string) [][]rune {
	m := [][]rune{}
	for _, line := range strings.Split(s, "\n") {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		m = append(m, row)
	}
	return m
}

type coordinate struct {
	x, y int
}

func asteroids(astr [][]rune) (asteroids []coordinate) {
	for y, row := range astr {
		for x, col := range row {
			if col == '#' {
				asteroids = append(asteroids, coordinate{x, y})
			}
		}
	}
	return
}

func countVisible(asteroids []coordinate) (int, coordinate) {
	var max int
	var position coordinate
	for _, a := range asteroids {
		memo := make(map[float64][]coordinate)
		for _, b := range asteroids {
			if a != b {
				angle := -math.Atan2(float64(b.x-a.x), float64(b.y-a.y))
				memo[angle] = append(memo[angle], b)
			}
		}
		if len(memo) > max {
			max = len(memo)
			position = a
		}
	}
	return max, position
}
