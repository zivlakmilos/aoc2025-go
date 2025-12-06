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
	input := getInput()

	lines := strings.Split(input, "\n")
	operations := lines[len(lines)-1]

	var op byte
	res := 0
	opres := 0

	x := 0
	for {
		found := false
		if x < len(operations) && operations[x] != ' ' {
			op = operations[x]
			res += opres
			if op == '*' {
				opres = 1
			} else {
				opres = 0
			}
		}
		current := 0
		for y := range len(lines) - 1 {
			if x < len(lines[y]) {
				found = true
				if lines[y][x] != ' ' {
					current *= 10
					current += int(lines[y][x] - '0')
				}
			}
		}
		if current > 0 {
			if op == '*' {
				opres *= current
			} else {
				opres += current
			}
		}
		x++

		if !found {
			res += opres
			break
		}
	}

	fmt.Printf("Grand total: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
