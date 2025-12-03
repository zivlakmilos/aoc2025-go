package main

import (
	"fmt"
	"strconv"
	"strings"
)

func keepLargest(s string, keep int) string {
	n := len(s)
	k := n - keep
	stack := []byte{}

	for i := range n {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < s[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, s[i])
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	return string(stack)
}

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
	input := getInput()
	res := uint64(0)

	for line := range strings.SplitSeq(input, "\n") {
		line = keepLargest(line, 12)

		joltage, _ := strconv.ParseUint(line, 10, 64)
		res += joltage
	}

	fmt.Printf("Total output joltage is: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
