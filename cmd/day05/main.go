package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([][2]uint64, []uint64) {
	var db [][2]uint64
	var ids []uint64

	inputs := strings.Split(input, "\n\n")

	for line := range strings.SplitSeq(inputs[0], "\n") {
		numbers := strings.Split(line, "-")
		num1, _ := strconv.ParseUint(numbers[0], 10, 64)
		num2, _ := strconv.ParseUint(numbers[1], 10, 64)
		db = append(db, [2]uint64{num1, num2})
	}

	for line := range strings.SplitSeq(inputs[1], "\n") {
		num, _ := strconv.ParseUint(line, 10, 64)
		ids = append(ids, num)
	}

	return db, ids
}

func solvePuzzle01() {
	input := getInput()
	db, ids := parseInput(input)
	res := 0

	for _, id := range ids {
		for _, r := range db {
			if id >= r[0] && id <= r[1] {
				res++
				break
			}
		}
	}

	fmt.Printf("Fresh ingredients: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
