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
}

func main() {
	solvePzzle01()
	solvePzzle02()
}
