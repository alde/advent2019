package day10

import (
	"math"
	"sort"
	"strings"

	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := lib.ReadFile("day10/input.txt")
	amap := parse(input)
	asteroids := asteroids(amap)

	seen, pos := countVisible(asteroids)

	logrus.WithFields(logrus.Fields{
		"x":       pos.x,
		"y":       pos.y,
		"visible": len(seen),
	}).Info("best position for base")

	destroyed := destroy(seen, pos)
	logrus.WithFields(logrus.Fields{
		"200th destroyed": destroyed[199].x*100 + destroyed[199].y,
	}).Info("after the laser")

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

func countVisible(asteroids []coordinate) (map[float64][]coordinate, coordinate) {
	seen := make(map[float64][]coordinate)
	position := coordinate{}

	for _, a := range asteroids {
		memo := make(map[float64][]coordinate)
		for _, b := range asteroids {
			if a != b {
				angle := -math.Atan2(float64(b.x-a.x), float64(b.y-a.y))
				memo[angle] = append(memo[angle], b)
			}
		}
		if len(memo) > len(seen) {
			seen = memo
			position = a
		}
	}

	return seen, position
}

func destroy(roids map[float64][]coordinate, base coordinate) []coordinate {
	for a := range roids {
		sort.Slice(roids[a], func(i, j int) bool {
			return dist(base, roids[a][i]) > dist(base, roids[a][j])
		})

		for len(roids[a]) > 1 {
			i := a + 2*math.Pi*float64(len(roids[a])-1)
			roids[i] = append(roids[i], roids[a][0])
			roids[a] = roids[a][1:]
		}
	}
	angles := []float64{}
	for a := range roids {
		angles = append(angles, a)
	}
	memo := []coordinate{}
	sort.Float64s(angles)
	for x := range angles {
		memo = append(memo, roids[angles[x]][0])
	}
	return memo
}

func dist(a, b coordinate) float64 {
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}
