package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x   int
	y   int
	z   int
	raw string
}

type Dist struct {
	a    JunctionBox
	b    JunctionBox
	dist int
}

func parseInput(input string) []JunctionBox {
	var res []JunctionBox

	for line := range strings.SplitSeq(input, "\n") {
		cols := strings.Split(line, ",")
		x, _ := strconv.Atoi(cols[0])
		y, _ := strconv.Atoi(cols[1])
		z, _ := strconv.Atoi(cols[2])
		res = append(res, JunctionBox{
			x:   x,
			y:   y,
			z:   z,
			raw: line,
		})
	}

	return res
}

func calcDist(a, b JunctionBox) int {
	x := a.x - b.x
	y := a.y - b.y
	z := a.z - b.z

	return x*x + y*y + z*z
}

func find(parent map[string]string, x string) string {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent map[string]string, a, b string) {
	rootA := find(parent, a)
	rootB := find(parent, b)
	if rootA != rootB {
		parent[rootA] = rootB
	}
}

func solvePuzzle01() {
	input := getInput()
	pos := parseInput(input)

	dist := make([]Dist, 0, len(pos)*len(pos))

	for i := range pos {
		for j := i + 1; j < len(pos); j++ {
			d := calcDist(pos[i], pos[j])
			dist = append(dist, Dist{
				a:    pos[i],
				b:    pos[j],
				dist: d,
			})
		}
	}

	slices.SortFunc(dist, func(a, b Dist) int {
		return a.dist - b.dist
	})
	dist = dist[:1000]

	parent := make(map[string]string)
	for _, p := range pos {
		parent[p.raw] = p.raw
	}

	for _, d := range dist {
		union(parent, d.a.raw, d.b.raw)
	}

	sizes := make(map[string]int)
	for _, p := range pos {
		root := find(parent, p.raw)
		sizes[root]++
	}

	res := make([]int, 0)
	for _, size := range sizes {
		res = append(res, size)
	}

	slices.SortFunc(res, func(a, b int) int {
		return b - a
	})

	for len(res) < 3 {
		res = append(res, 1)
	}

	fmt.Printf("res: %d\n", res[0]*res[1]*res[2])
}

func solvePuzzle02() {
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
