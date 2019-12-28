package day9

import (
	"alde.nu/advent/lib"
	"github.com/sirupsen/logrus"
)

// Run todays challenge
func Run() {
	input := lib.ReadFileAsIntSlice("day9/input.txt")
	res1 := boost(input, 1)
	logrus.WithField("keycode", res1).Info("BOOST code")
	res2 := boost(input, 2)
	logrus.WithField("coordinates", res2).Info("BOOST code")
}

func boost(input []int, seed int) int {
	output := make(chan int)
	ichan := make(chan int)
	halt := make(chan []int)
	go lib.OpCode(input, ichan, output, halt)
	ichan <- seed
	for {
		select {
		case o := <-output:
			return o
		case <-halt:
			return -1
		}
	}
}
