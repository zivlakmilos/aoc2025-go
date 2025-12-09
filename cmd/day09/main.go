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
	w := a.x - b.x
	h := a.y - b.y

	if w < 0 {
		w = -w
	}
	if h < 0 {
		h = -h
	}
	w++
	h++

	return w * h
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func pointOnPolygonEdge(p Point, polygon []Point) bool {
	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]

		if p1.x == p2.x && p.x == p1.x {
			if p.y >= min(p1.y, p2.y) && p.y <= max(p1.y, p2.y) {
				return true
			}
		} else if p1.y == p2.y && p.y == p1.y {
			if p.x >= min(p1.x, p2.x) && p.x <= max(p1.x, p2.x) {
				return true
			}
		} else {
			cross := (p2.x-p1.x)*(p.y-p1.y) - (p2.y-p1.y)*(p.x-p1.x)
			if cross == 0 {
				if p.x >= min(p1.x, p2.x) && p.x <= max(p1.x, p2.x) &&
					p.y >= min(p1.y, p2.y) && p.y <= max(p1.y, p2.y) {
					return true
				}
			}
		}
	}
	return false
}

func pointInPolygon(p Point, polygon []Point) bool {
	if pointOnPolygonEdge(p, polygon) {
		return true
	}

	if len(polygon) < 3 {
		return false
	}

	inside := false
	j := len(polygon) - 1

	for i := range polygon {
		p1 := polygon[i]
		p2 := polygon[j]

		if p1.y != p2.y {
			if (p1.y > p.y) != (p2.y > p.y) {
				xIntersect := float64(p1.x) + float64(p2.x-p1.x)*float64(p.y-p1.y)/float64(p2.y-p1.y)
				if float64(p.x) < xIntersect {
					inside = !inside
				}
			}
		}
		j = i
	}

	return inside
}

func crosses(a1, a2, b1, b2 Point) bool {
	ccw := func(p1, p2, p3 Point) int {
		val := (p2.x-p1.x)*(p3.y-p1.y) - (p2.y-p1.y)*(p3.x-p1.x)
		if val > 0 {
			return 1
		}
		if val < 0 {
			return -1
		}
		return 0
	}

	d1 := ccw(b1, b2, a1)
	d2 := ccw(b1, b2, a2)
	d3 := ccw(a1, a2, b1)
	d4 := ccw(a1, a2, b2)

	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

func polygonEdgeCrossesRectangle(rectP1, rectP2, segP1, segP2 Point) bool {
	minX := min(rectP1.x, rectP2.x)
	maxX := max(rectP1.x, rectP2.x)
	minY := min(rectP1.y, rectP2.y)
	maxY := max(rectP1.y, rectP2.y)

	if (segP1.x < minX && segP2.x < minX) || (segP1.x > maxX && segP2.x > maxX) {
		return false
	}
	if (segP1.y < minY && segP2.y < minY) || (segP1.y > maxY && segP2.y > maxY) {
		return false
	}

	onBoundary := func(p Point) bool {
		return ((p.x == minX || p.x == maxX) && p.y >= minY && p.y <= maxY) ||
			((p.y == minY || p.y == maxY) && p.x >= minX && p.x <= maxX)
	}

	if onBoundary(segP1) && onBoundary(segP2) {
		if (segP1.x == segP2.x && (segP1.x == minX || segP1.x == maxX)) ||
			(segP1.y == segP2.y && (segP1.y == minY || segP1.y == maxY)) {
			return false
		}
	}

	rectEdges := [][]Point{
		{{minX, minY}, {maxX, minY}},
		{{maxX, minY}, {maxX, maxY}},
		{{maxX, maxY}, {minX, maxY}},
		{{minX, maxY}, {minX, minY}},
	}

	for _, edge := range rectEdges {
		if crosses(segP1, segP2, edge[0], edge[1]) {
			return true
		}
	}

	return false
}

func rectangleInsidePolygon(p1, p2 Point, polygon []Point) bool {
	if p1.x == p2.x || p1.y == p2.y {
		return false
	}

	corners := []Point{
		{p1.x, p1.y},
		{p1.x, p2.y},
		{p2.x, p1.y},
		{p2.x, p2.y},
	}

	for _, corner := range corners {
		if !pointInPolygon(corner, polygon) {
			return false
		}
	}

	for i := range polygon {
		next := (i + 1) % len(polygon)
		if polygonEdgeCrossesRectangle(p1, p2, polygon[i], polygon[next]) {
			return false
		}
	}

	return true
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
	input := getInput()
	points := parseInput(input)

	res := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if rectangleInsidePolygon(points[i], points[j], points) {
				rectArea := calcArea(points[i], points[j])
				if rectArea > res {
					res = rectArea
				}
			}
		}
	}

	fmt.Printf("Largest area: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
