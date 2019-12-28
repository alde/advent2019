package lib

import (
	"github.com/sirupsen/logrus"
)

// OpCodes
const (
	Add                = 1
	Multiply           = 2
	Read               = 3
	Write              = 4
	JumpIfTrue         = 5
	JumpIfFalse        = 6
	LessThan           = 7
	Equal              = 8
	AdjustRelativeBase = 9
	Halt               = 99
)

// OpCode implements the computer
func OpCode(program []int, input <-chan int, output chan<- int, halt chan<- []int) {
	memory := make([]int, len(program))
	copy(memory, program)

	relbase := 0
	ip := 0
	for {
		instruction := memory[ip]
		opcode := instruction % 100

		arg := func(offset int) (addr int) {
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0:
				addr = memory[ip+offset]
			case 1:
				addr = ip + offset
			case 2:
				addr = relbase + memory[ip+offset]
			}
			if len(memory) <= addr {
				memory = append(memory, make([]int, addr-len(memory)+1)...)
			}
			return
		}

		switch opcode {
		case Add: // 1
			memory[arg(3)] = memory[arg(1)] + memory[arg(2)]
			ip += 4

		case Multiply: // 2
			memory[arg(3)] = memory[arg(1)] * memory[arg(2)]
			ip += 4

		case Read: // 3
			memory[arg(1)] = <-input
			ip += 2

		case Write: // 4
			output <- memory[arg(1)]
			ip += 2

		case JumpIfTrue: // 5
			if memory[arg(1)] != 0 {
				ip = memory[arg(2)]
				continue
			}
			ip += 3

		case JumpIfFalse: // 6
			if memory[arg(1)] == 0 {
				ip = memory[arg(2)]
				continue
			}
			ip += 3

		case LessThan: // 7
			if memory[arg(1)] < memory[arg(2)] {
				memory[arg(3)] = 1
			} else {
				memory[arg(3)] = 0
			}
			ip += 4

		case Equal: // 8
			if memory[arg(1)] == memory[arg(2)] {
				memory[arg(3)] = 1
			} else {
				memory[arg(3)] = 0
			}
			ip += 4

		case AdjustRelativeBase: // 9
			relbase += memory[arg(1)]
			ip += 2

		case 99:
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
