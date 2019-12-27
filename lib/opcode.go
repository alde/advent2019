package lib

import (
	"github.com/sirupsen/logrus"
)

// OpCodes
const (
	Add         = 1
	Multiply    = 2
	Read        = 3
	Write       = 4
	JumpIfTrue  = 5
	JumpIfFalse = 6
	LessThan    = 7
	Equal       = 8
	Halt        = 99
)

// OpCode implements the computer
func OpCode(program []int, input <-chan int, output chan<- int, halt chan<- []int) {
	memory := make([]int, len(program))
	copy(memory, program)

	ip := 0
	for {
		instruction := memory[ip]
		opcode := instruction % 100

		val := func(offset int) int {
			parameter := memory[offset+ip]
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0: // position mode
				return memory[parameter]
			case 1: // immediate mode
				return parameter
			default:
				logrus.WithField("mode", mode).Fatal("invalid mode")
				return -1
			}
		}

		switch opcode {
		case Add:
			a, b, c := val(1), val(2), memory[ip+3]
			memory[c] = a + b
			ip += 4

		case Multiply:
			a, b, c := val(1), val(2), memory[ip+3]
			memory[c] = a * b
			ip += 4

		case Read:
			a := memory[ip+1]
			memory[a] = <-input
			ip += 2

		case Write:
			a := val(1)
			output <- a
			ip += 2

		case JumpIfTrue:
			a, b := val(1), val(2)
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}

		case JumpIfFalse:
			a, b := val(1), val(2)
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}

		case LessThan:
			a, b, c := val(1), val(2), memory[ip+3]
			if a < b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case Equal:
			a, b, c := val(1), val(2), memory[ip+3]
			if a == b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case Halt:
			halt <- memory
			return

		default:
			logrus.WithField("opcode", opcode).Fatal("unexpected opcode")
			return
		}
	}
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
