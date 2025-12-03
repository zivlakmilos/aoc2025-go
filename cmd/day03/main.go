package main

import (
	"fmt"
	"strings"
)

func solvePuzzle01() {
	input := getInput()
	res := 0

	for line := range strings.SplitSeq(input, "\n") {
		digit1Idx := 0
		digit2Idx := 0

		for idx := range len(line) - 1 {
			if line[idx] > line[digit1Idx] {
				digit1Idx = idx
				if line[idx] == '9' {
					break
				}
			}
		}

		digit2Idx = digit1Idx + 1
		for idx := digit1Idx + 2; idx < len(line); idx++ {
			if line[idx] > line[digit2Idx] {
				digit2Idx = idx
				if line[idx] == '9' {
					break
				}
			}
		}

		joltage := (line[digit1Idx]-'0')*10 + (line[digit2Idx] - '0')
		res += int(joltage)
	}

	fmt.Printf("Total output joltage is: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
