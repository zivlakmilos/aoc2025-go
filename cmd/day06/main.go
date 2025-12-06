package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePuzzle01() {
	input := getInput()

	lines := strings.Split(input, "\n")

	op := strings.Fields(lines[len(lines)-1])
	res := make([]int, len(op))

	for i, o := range op {
		if o == "*" {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}

	for i := range len(lines) - 1 {
		col := strings.Fields(lines[i])
		for j := range col {
			if op[j] == "*" {
				val, _ := strconv.Atoi(col[j])
				res[j] *= val
			} else {
				val, _ := strconv.Atoi(col[j])
				res[j] += val
			}
		}
	}

	total := 0
	for _, r := range res {
		total += r
	}

	fmt.Printf("Grand total: %d\n", total)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
