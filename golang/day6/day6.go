package day6

import (
	"strings"

	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := lib.ReadFile("day6/input.txt")
	planets := parse(input)
	logrus.Info("Day 6")
	logrus.WithField("orbits", planets.countAllOrbits()).Info("total orbits")

	logrus.WithField("transfers", planets.getDistance("YOU", "SAN")).Info("transfers to reach SAN")
}

func parse(input string) Planets {
	planets := Planets{}

	for _, s := range strings.Split(input, "\n") {
		r := strings.Split(s, ")")
		planets[r[1]] = Planet{
			Name:   r[0],
			Orbits: 0,
		}
	}
	return planets
}

// Planets map
type Planets map[string]Planet

// Planet struct
type Planet struct {
	Name   string
	Orbits int
}

func (p Planets) countOrbits(name string) int {
	planet, ok := p[name]
	if !ok {
		return 0
	}
	if planet.Orbits == 0 {
		planet.Orbits = 1 + p.countOrbits(planet.Name)
		p[name] = planet
	}
	return planet.Orbits
}

func (p Planets) countAllOrbits() int {
	count := 0
	for name := range p {
		count += p.countOrbits(name)
	}
	return count
}

func (p Planets) getPath(from string) []string {
	path := []string{}
	planet, ok := p[from]
	for ok {
		path = append(path, planet.Name)
		planet, ok = p[planet.Name]
	}
	return path
}

func (p Planets) getDistance(from string, to string) int {
	pathFrom := p.getPath(from)
	pathTo := p.getPath(to)
	path := travelFrom(pathFrom, pathTo, []string{})
	return len(path)
}

func travelFrom(pathFrom []string, pathTo []string, path []string) []string {
	planet := pathFrom[0]
	tail := pathFrom[1:]
	idx := index(pathTo, planet)
	if idx != -1 {
		return travelTo(pathTo, idx, path)
	}
	return travelFrom(tail, pathTo, append(path, planet))
}

func travelTo(pathTo []string, index int, path []string) []string {
	if index == 0 {
		return path
	}
	planet := pathTo[index]
	return travelTo(pathTo, index-1, append(path, planet))
}

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
