package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(str string) int {
	res, _ := strconv.Atoi(str[1:])

	if str[0] == 'L' {
		res = -res
	}

	return res
}

func solvePuzzle01() {
	res := 0
	dial := 50

	input := getInput()
	for line := range strings.SplitSeq(input, "\n") {
		rot := parseLine(line)
		dial += rot
		for dial < 0 {
			dial = 100 + dial
		}

		dial %= 100
		if dial == 0 {
			res++
		}
	}

	fmt.Printf("Password to open door is: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
