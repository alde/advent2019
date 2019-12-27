package day3

import (
	"sort"
	"strconv"
	"strings"

	"alde.nu/advent/lib"

	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.Info("Day 3")
	input := parse(lib.ReadFile("day3/input"))

	result1 := manhattan(input)
	logrus.WithField("distance", result1).Info("manhattan")
}

func drawLine(input []string, out chan<- []point) {
	line := []point{}
	stepCount := 0
	x := 0
	y := 0
	for _, step := range input {
		stepCount++
		direction := step[0]
		steps, _ := strconv.Atoi(step[1:])
		for s := 0; s < steps; s++ {
			switch direction {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			case 'D':
				y--
			}
			line = append(line, point{x: x, y: y})
		}
	}
	out <- line
}

func manhattan(input [][]string) int {
	line1 := make(chan []point)
	line2 := make(chan []point)
	go drawLine(input[0], line1)
	go drawLine(input[1], line2)
	var o1, o2 []point
loop:
	for {
		select {
		case o1 = <-line1:
			if o2 != nil {
				break loop
			}
		case o2 = <-line2:
			if o1 != nil {
				break loop
			}
		}
	}
	intersection := intersect(o1, o2)
	sort.Slice(intersection, func(i, j int) bool {
		return intersection[i].distance() < intersection[j].distance()
	})
	return intersection[0].distance()
}

type point struct {
	x int
	y int
}

func (p point) distance() int {
	return abs(p.x) + abs(p.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func parse(input string) [][]string {
	i := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		l := []string{}
		for _, direction := range strings.Split(line, ",") {
			l = append(l, direction)
		}
		i = append(i, l)
	}
	return i
}

func intersect(a []point, b []point) []point {
	set := []point{}
	hash := make(map[point]bool)

	for i := 0; i < len(a); i++ {
		hash[a[i]] = true
	}

	for i := 0; i < len(b); i++ {
		el := b[i]
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}
