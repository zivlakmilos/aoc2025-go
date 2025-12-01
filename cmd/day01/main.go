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
	res := 0
	dial := 50

	input := getInput()
	for line := range strings.SplitSeq(input, "\n") {
		rot := parseLine(line)

		isLeft := rot < 0
		if isLeft {
			rot = -rot
		}

		for range rot {
			if isLeft {
				dial--
				if dial == 0 {
					res++
				} else if dial < 0 {
					dial = 99
				}
			} else {
				dial++
				for dial >= 100 {
					dial = 0
					res++
				}
			}
		}
	}

	fmt.Printf("Password to open door is: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
