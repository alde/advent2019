package main

import "fmt"

func main() {
	count1 := partOne(256310, 732736)
	fmt.Printf("%d passcodes in the given range\n", count1)

	count2 := partTwo(256310, 732736)
	fmt.Printf("%d passcodes in the given range with new constraints\n", count2)
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

// Part One
func partOne(from, to int) int {
	return innerPartOne(from, to, 0)
}

func innerPartOne(from, to, validCount int) int {
	if from > to {
		return validCount
	}

	if partOneValidator(from) {
		return innerPartOne(from+1, to, validCount+1)
	}
	return innerPartOne(from+1, to, validCount)
}

func partOneValidator(candidate int) bool {
	list := intToSlice(candidate, []int{})
	head, tail := list[0], list[1:]
	return partOneValidatorInner(head, tail, false, false)
}

func partOneValidatorInner(head int, list []int, repeat bool, valid bool) bool {
	if len(list) == 0 {
		return repeat && valid
	}
	h0, tail := list[0], list[1:]
	if h0 > head {
		return partOneValidatorInner(h0, tail, repeat, true)
	}
	if h0 == head {
		return partOneValidatorInner(h0, tail, true, true)
	}
	if h0 < head {
		return false
	}
	return valid
}

// Part 2
func partTwo(from, to int) int {
	return innerPartTwo(from, to, 0)
}

func innerPartTwo(from, to, validCount int) int {
	if from > to {
		return validCount
	}

	if partTwoValidator(from) {
		return innerPartTwo(from+1, to, validCount+1)
	}
	return innerPartTwo(from+1, to, validCount)
}


func partTwoValidator(candidate int) bool {
	list := intToSlice(candidate, []int{})
	head, tail := list[0], list[1:]
	memo := make(map[int]int)
	return partTwoValidatorInner(head, tail, memo, false)
}

func partTwoValidatorInner(head int, list []int, memo map[int]int, valid bool) bool {
	if len(list) == 0 {
		return validateMemo(memo)
	}
	newHead, tail := list[0], list[1:]
	if newHead > head {
		return partTwoValidatorInner(newHead, tail, memo, true)
	}
	if newHead == head {
		if _, ok := memo[newHead] ; !ok {
			memo[newHead] = 1
		}
		memo[newHead]++
		return partTwoValidatorInner(newHead, tail, memo, true)
	}
	if newHead < head {
		return false
	}
	return valid
}

func validateMemo(memo map[int]int) bool {
	for _, v := range memo {
		if v == 2 {
			return true
		}
	}
	return false
}