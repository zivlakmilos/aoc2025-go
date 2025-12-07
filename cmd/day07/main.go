package main

import (
	"fmt"
	"strings"
)

func parseInput(input string) [][]byte {
	var res [][]byte

	for line := range strings.SplitSeq(input, "\n") {
		var row []byte
		for i := range line {
			row = append(row, line[i])
		}
		res = append(res, row)
	}

	return res
}

func solveTimeline(data [][]byte, mem map[uint64]int, row, col int) int {
	if row >= len(data) {
		return 1
	}
	key := (uint64(row) << 32) | uint64(col)
	if val, ok := mem[key]; ok {
		return val
	}

	res := 0

	if data[row][col] == '^' {
		res += solveTimeline(data, mem, row+1, col+1)
		res += solveTimeline(data, mem, row+1, col-1)
	} else {
		res += solveTimeline(data, mem, row+1, col)
	}

	mem[key] = res

	return res
}

func solvePzzle01() {
	input := getInput()
	data := parseInput(input)

	res := 0
	for i := 1; i < len(data); i++ {
		for j := range len(data[i]) {
			if data[i-1][j] == '|' || data[i-1][j] == 'S' {
				if data[i][j] == '^' {
					data[i][j+1] = '|'
					data[i][j-1] = '|'
					res++
				} else {
					data[i][j] = '|'
				}
			}
		}
	}

	fmt.Printf("Total splits: %d\n", res)
}

func solvePzzle02() {
	input := getInput()
	data := parseInput(input)

	col := 0
	for i := range len(data[0]) {
		if data[0][i] == 'S' {
			col = i
			break
		}
	}

	mem := map[uint64]int{}
	res := solveTimeline(data, mem, 0, col)

	fmt.Printf("Total timelines: %d\n", res)
}

func main() {
	solvePzzle01()
	solvePzzle02()
}
