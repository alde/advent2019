package day3

import (
	"sort"
	"strconv"
	"strings"

	"alde.nu/advent/lib"

	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := parse(lib.ReadFile("day3/input"))

	result1, result2 := manhattan(input)
	logrus.WithField("distance", result1).Info("manhattan")
	logrus.WithField("steps", result2).Info("fewest combined steps")
}

func drawLine(input []string, out chan<- res) {
	line := []point{}
	stepsToPos := make(map[point]int)
	stepCount := 0
	x := 0
	y := 0
	for _, step := range input {
		direction := step[0]
		steps, _ := strconv.Atoi(step[1:])
		for s := 0; s < steps; s++ {
			stepCount++
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
			p := point{x: x, y: y}
			if _, ok := stepsToPos[p]; !ok {
				stepsToPos[p] = stepCount
			}
			line = append(line, p)
		}
	}
	out <- res{line, stepsToPos}
}

type res struct {
	line  []point
	steps map[point]int
}

func manhattan(input [][]string) (int, int) {
	res1 := make(chan res)
	res2 := make(chan res)
	go drawLine(input[0], res1)
	go drawLine(input[1], res2)
	var o1, o2 res
loop:
	for {
		select {
		case o1 = <-res1:
			if o2.line != nil {
				break loop
			}
		case o2 = <-res2:
			if o1.line != nil {
				break loop
			}
		}
	}
	intersection := intersect(o1.line, o2.line)
	var i2 []point
	copy(i2, intersection)
	sort.Slice(intersection, func(i, j int) bool {
		return intersection[i].distance() < intersection[j].distance()
	})
	var shortest int
	for _, p := range intersection {
		current := o1.steps[p] + o2.steps[p]
		if shortest == 0 || current < shortest {
			shortest = current
		}
	}

	manhattan := intersection[0].distance()

	return manhattan, shortest
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
