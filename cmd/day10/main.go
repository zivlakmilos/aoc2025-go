package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	Goal     int
	Lights   int
	Buttons  []int
	Joltages []int
}

func parseInput(input string) []Machine {
	var res []Machine

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, " ")

		goal := 0
		lights := 0
		buttons := make([]int, 0)
		joltages := make([]int, 0)

		for _, part := range parts {
			inner := part[1 : len(part)-1]

			switch part[0] {
			case '[':
				lights = len(inner)
				for j, c := range inner {
					if c == '#' {
						goal |= (1 << j)
					}
				}
			case '(':
				button := 0
				for _, ns := range strings.Split(inner, ",") {
					num, _ := strconv.Atoi(ns)
					button |= (1 << num)
				}
				buttons = append(buttons, button)
			case '{':
				for _, ns := range strings.Split(inner, ",") {
					num, _ := strconv.Atoi(ns)
					joltages = append(joltages, num)
				}
			}
		}

		res = append(res,
			Machine{
				Goal:     goal,
				Lights:   lights,
				Buttons:  buttons,
				Joltages: joltages,
			})
	}

	return res
}

func minButtonPresses(m Machine) int {
	numButtons := len(m.Buttons)
	minPresses := math.MaxInt

	for mask := range 1 << numButtons {
		lights := 0
		presses := 0

		for i := range numButtons {
			if mask&(1<<i) != 0 {
				lights ^= m.Buttons[i]
				presses++
			}
		}

		if lights == m.Goal && presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}

func solvePuzzle01() {
	input := getInput()
	data := parseInput(input)
	res := 0

	for _, d := range data {
		res += minButtonPresses(d)
	}
	fmt.Printf("result: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
