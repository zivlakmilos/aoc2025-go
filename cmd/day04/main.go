package main

import (
	"fmt"
	"strings"
)

func parseInput(input string) [][]byte {
	res := [][]byte{}

	for line := range strings.SplitSeq(input, "\n") {
		row := make([]byte, len(line))
		for i := range line {
			if line[i] == '@' {
				row[i] = 1
			} else {
				row[i] = 0
			}
		}
		res = append(res, row)
	}

	return res
}

func getElement(matrix [][]byte, x, y int) byte {
	if y < 0 {
		return 0
	}
	if x < 0 {
		return 0
	}
	if y >= len(matrix) {
		return 0
	}
	if x >= len(matrix[y]) {
		return 0
	}

	return matrix[y][x]
}

func countNeighbours(matrix [][]byte, x, y int) int {
	res := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			res += int(getElement(matrix, i, j))
		}
	}

	res -= int(getElement(matrix, x, y))

	return res
}

func solvePuzzle01() {
	input := getInput()
	matrix := parseInput(input)
	res := 0

	for y := range matrix {
		for x := range matrix[y] {
			if getElement(matrix, x, y) > 0 && countNeighbours(matrix, x, y) < 4 {
				res++
			}
		}
	}

	fmt.Printf("Rolls of paper: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
