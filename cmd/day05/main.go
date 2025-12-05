package main

import (
	"fmt"
	"slices"
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
	input := getInput()
	db, _ := parseInput(input)
	res := uint64(0)

	slices.SortFunc(db, func(a, b [2]uint64) int {
		return int(a[0] - b[0])
	})

	merged := [][2]uint64{db[0]}
	count := 0
	for i := 1; i < len(db); i++ {
		if db[i][0] <= merged[count][1]+1 {
			if db[i][1] > merged[count][1] {
				merged[count][1] = db[i][1]
			}
		} else {
			merged = append(merged, db[i])
			count++
		}
	}

	for _, r := range merged {
		res += r[1] - r[0] + 1
	}

	fmt.Printf("Fresh ingredients: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
