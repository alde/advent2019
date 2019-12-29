package main

import (
	"flag"
	"fmt"

	"alde.nu/advent/day1"
	"alde.nu/advent/day2"
	"alde.nu/advent/day3"
	"alde.nu/advent/day4"
	"alde.nu/advent/day5"
	"alde.nu/advent/day6"
	"alde.nu/advent/day7"
	"alde.nu/advent/day8"
	"alde.nu/advent/day9"
	"alde.nu/advent/day10"
	"github.com/sirupsen/logrus"
)

func main() {
	solutions := map[int]func(){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
		5: day5.Run,
		6: day6.Run,
		7: day7.Run,
		8: day8.Run,
		9: day9.Run,
		10: day10.Run,
	}
	execute := func(day int) {
		logrus.Infof("Day %d\n", day)
		solutions[day]()
	}
	var d int
	flag.IntVar(&d, "day", 0, fmt.Sprintf("specify day to run [1..%d]", len(solutions)))
	flag.Parse()

	if d == 0 {
		for day := range solutions {
			execute(day)
		}
	} else {
		execute(d)
	}
}
