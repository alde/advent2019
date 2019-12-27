package day2

import (
	"strconv"
	"strings"

	"alde.nu/advent/lib"

	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.Info("Day 2")
	input := parse(lib.ReadFile("day2/input"))

	result1 := gravityAssist(input)
	logrus.WithField("output", result1).Info("signal")

	result2 := findInputs(input, 19690720)
	logrus.WithFields(logrus.Fields{
		"inputs": result2,
		"code": 100 * result2.noun + result2.verb,
	}).Info("inputs")
}

func parse(in string) []int {
	out := []int{}
	for _, s := range strings.Split(in, ",") {
		i, _ := strconv.Atoi(s)
		out = append(out, i)
	}
	return out
}

type inputs struct {
	noun int
	verb int
}
func findInputs(input []int, targetSignal int) *inputs {
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			newCode := append([]int{input[0], noun, verb}, input[3:]...)
			if gravityAssist(newCode) == targetSignal {
				return &inputs{
					noun: noun,
					verb: verb,
				}
			}
		}
	}
	return nil
}

func gravityAssist(input []int) int {
	out := make(chan []int)
	go lib.OpCode(input, nil, nil, out)
	select {
	case res := <- out:
		return res[0]
	}
	logrus.Fatal("should never get here")
	return -1
}
