package day8

import (
	"fmt"
	"strconv"

	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := lib.ReadFile("day8/input.txt")
	layers := parseLayers(input, 25, 6)
	logrus.Info("Day 8")
	checksum := checksum(layers)
	logrus.WithField("checksum", checksum).Info("checksum for image")
	flattened := flatten(layers, 25*6)
	logrus.Info("final image")
	print(flattened, 25)
}

func process(input string, width int, height int, collector [][]int) [][]int {
	if len(input) == 0 {
		return collector
	}
	layer := []int{}
	offset := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			d, _ := strconv.Atoi(string(input[offset]))
			layer = append(layer, d)
			offset++
		}
	}
	collector = append(collector, layer)
	return process(input[offset:], width, height, collector)
}

func parseLayers(input string, width int, height int) [][]int {
	layers := process(input, width, height, [][]int{})
	return layers
}

func checksum(layers [][]int) int {
	fewestZeroLayer := []int{}
	fewestZeroCount := 0
	for _, layer := range layers {
		c := count(layer, 0)
		if c < fewestZeroCount || fewestZeroCount == 0 {
			fewestZeroLayer = layer
			fewestZeroCount = c
		}
	}
	return count(fewestZeroLayer, 1) * count(fewestZeroLayer, 2)
}

func count(layer []int, candidate int) int {
	count := 0
	for _, i := range layer {
		if i == candidate {
			count++
		}
	}
	return count
}

func flatten(layers [][]int, size int) []int {
	final := make([]int, size)
	for i := 0; i < size; i++ {
		final[i] = 2
	}
	for _, layer := range layers {
		for idx, pixel := range layer {
			if final[idx] >= 2 {
				final[idx] = pixel
			}
		}
	}

	return final
}

func print(data []int, rowLength int) {
	for i, s := range data {
		if i%rowLength == 0 && i > 0 {
			fmt.Println()
		}
		if s == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("█")
		}
	}
	fmt.Println()
}
