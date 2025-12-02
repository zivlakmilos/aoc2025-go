package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func parseInput(input string) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for r := range strings.SplitSeq(input, ",") {
			values := strings.Split(r, "-")
			from, _ := strconv.Atoi(values[0])
			to, _ := strconv.Atoi(values[1])
			if !yield(from, to) {
				return
			}
		}
	}
}

func checkIsValid01(id int) bool {
	digits := make([]uint8, 100)
	i := 0

	for id > 0 {
		dig := uint8(id % 10)
		id /= 10

		digits[i] = dig
		i++
	}

	if i%2 != 0 {
		return true
	}

	for j := range i / 2 {
		if digits[j] != digits[i/2+j] {
			return true
		}
	}

	return false
}

func solvePuzzle01() {
	input := getInput()
	res := 0

	for from, to := range parseInput(input) {
		for i := from; i <= to; i++ {
			if !checkIsValid01(i) {
				res += i
			}
		}
	}

	fmt.Printf("Sum of invalid ids: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
