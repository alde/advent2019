package day7

import (
	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	logrus.Info("Day 7")
	input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 43, 64, 85, 98, 179, 260, 341, 422, 99999, 3, 9, 1001, 9, 3, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 1001, 9, 3, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 102, 3, 9, 9, 101, 4, 9, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}
	// part one
	logrus.WithField("signal", maxThrust(input, []int{0, 1, 2, 3, 4})).Info("max thrust")
	// part two
	logrus.WithField("signal", maxThrust(input, []int{5, 6, 7, 8, 9})).Info("max thrust with feedback loop")
}

func maxThrust(input []int, phases []int) int {
	max := 0
	for _, phase := range permutation(phases) {
		res := runComputer(input, phase)
		if res > max {
			max = res
		}
	}
	return max
}

func runComputer(input []int, phase []int) int {
	etoa := make(chan int, 1)
	atob := make(chan int)
	btoc := make(chan int)
	ctod := make(chan int)
	dtoe := make(chan int)

	halt := make(chan []int)

	// Start all amplifiers - they will wait for input
	go lib.OpCode(input, etoa, atob, halt)
	go lib.OpCode(input, atob, btoc, halt)
	go lib.OpCode(input, btoc, ctod, halt)
	go lib.OpCode(input, ctod, dtoe, halt)
	go lib.OpCode(input, dtoe, etoa, halt)

	// initialize phase settings
	etoa <- phase[0]
	atob <- phase[1]
	btoc <- phase[2]
	ctod <- phase[3]
	dtoe <- phase[4]

	// Send first signal
	etoa <- 0

	// Wait until the sequence ends
	for i := 0; i < 5; i++ {
		<-halt
	}

	// read the result.
	return <-etoa
}

func permutation(values []int) (result [][]int) {
	if len(values) == 1 {
		result = append(result, values)
		return result
	}
	for i, current := range values {
		others := make([]int, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, other := range permutation(others) {
			result = append(result, append(other, current))
		}
	}
	return result
}
