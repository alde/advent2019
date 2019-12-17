package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	planets := parse(strings.TrimSpace(string(input)))
	fmt.Printf("total orbits: %d\n", planets.countAllOrbits())
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

type Planets map[string]Planet

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
