package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func parseInput(input string) []Point {
	var res []Point

	for line := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		res = append(res, Point{
			x: x,
			y: y,
		})
	}

	return res
}

func calcArea(a, b Point) int {
	w := a.x - b.x + 1
	h := a.y - b.y + 1

	if w < 0 {
		w = -w
	}
	if h < 0 {
		h = -h
	}

	return w * h
}

func solvePuzzle01() {
	input := getInput()
	points := parseInput(input)

	res := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := calcArea(points[i], points[j])
			if area > res {
				res = area
			}
		}
	}

	fmt.Printf("Largest area: %d\n", res)
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
