package day05

import "C"
import (
	"AdventOfCode2021/util/aocmath"
	"AdventOfCode2021/util/aocstring"
	"log"
	"math"
	"regexp"
	"strings"
)

const (
	inputFile = "calendar/day05/input"

	DONT_INTERSECT = iota
	COLLINEAR
	DO_INTERSECT

	HORIZONTAL = "Horizontal"
	VERTICAL   = "Vertical"
	DIAGONAL   = "Diagonal"
)

var expr = regexp.MustCompile("\\d+")

type point struct {
	x int
	y int
}

type line struct {
	a point
	b point
}

func (l *line) direction() string {
	if l.a.y == l.b.y {
		return HORIZONTAL
	} else if l.a.x == l.b.x {
		return VERTICAL
	} else {
		return DIAGONAL
	}
}

func (l *line) minX() int {
	return aocmath.Min(l.a.x, l.b.x)
}

func (l *line) maxX() int {
	return aocmath.Max(l.a.x, l.b.x)
}

func (l *line) minY() int {
	return aocmath.Min(l.a.y, l.b.y)
}

func (l *line) maxY() int {
	return aocmath.Max(l.a.y, l.b.y)
}

func (l *line) withinLineBoundaries(p point) bool {
	xAxis := p.y >= aocmath.Min(l.a.x, l.b.x) && p.x <= aocmath.Max(l.a.x, l.b.y)
	yAxis := p.x >= aocmath.Min(l.a.x, l.b.x) && p.x <= aocmath.Max(l.a.x, l.b.y)

	return xAxis && yAxis
}

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
}

func part1(problem string) {
	log.Printf("Day 05 part 1: Hydrothermal Venture : %d", SolvePart1Second(problem))
}

func solvePart1(problem string) int {
	lines := ParseInput(problem)
	pointSet := make(map[point]struct{})

	filterOutDiagonals(&lines)

	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j {
				continue
			}
			intersectType, p := calcIntersectionPoint(line1, line2)

			switch intersectType {
			case COLLINEAR:
				addCollinearPoints(line1, line2, &pointSet)
				break
			case DO_INTERSECT:
				if _, ok := pointSet[p]; !ok {
					pointSet[p] = struct{}{}
				}
				break
			}
		}
	}

	return len(pointSet)
}

func addCollinearPoints(l1 line, l2 line, points *map[point]struct{}) {
	direction := l1.direction()

	switch direction {
	case HORIZONTAL:
		if l1.a.y != l2.a.y {
			return
		}
		maxmin := aocmath.Max(l1.minX(), l2.minX())
		minmax := aocmath.Min(l1.maxX(), l2.maxX())
		for x := maxmin; x <= minmax; x++ {
			(*points)[point{x, l1.a.y}] = struct{}{}
		}
		break
	case VERTICAL:
		if l1.a.x != l2.a.x {
			return
		}
		maxmin := aocmath.Max(l1.minY(), l2.minY())
		minmax := aocmath.Min(l1.maxY(), l2.maxY())
		for y := maxmin; y <= minmax; y++ {
			(*points)[point{l1.a.x, y}] = struct{}{}
		}
		break
	case DIAGONAL:
		break
	}
}

func filterOutDiagonals(lines *[]line) {
	for i := len(*lines) - 1; i >= 0; i-- {
		l := (*lines)[i]

		if l.a.x == l.b.x || l.a.y == l.b.y {
			continue
		}

		(*lines)[i] = (*lines)[len(*lines)-1]
		*lines = (*lines)[:len(*lines)-1]
	}
}

func calcIntersectionPoint(line1, line2 line) (int, point) {
	a1 := line1.b.y - line1.a.y
	b1 := line1.a.x - line1.b.x
	c1 := a1*(line1.a.x) + b1*(line1.a.y)

	// Line CD represented as a2x + b2y = c2
	a2 := line2.b.y - line2.a.y
	b2 := line2.a.x - line2.b.x
	c2 := a2*(line2.a.x) + b2*(line2.a.y)

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		// The lines are parallel. This is simplified
		// by returning a pair of FLT_MAX
		return COLLINEAR, point{-1, -1}
	} else if determinant < 0 {
		return DONT_INTERSECT, point{-1, -1}
	} else {
		x := int(float32(b2*c1-b1*c2) / float32(determinant))
		y := int(float32(a1*c2-a2*c1) / float32(determinant))

		if aocmath.Min(line1.a.x, line1.b.x) <= x && aocmath.Max(line1.a.x, line1.b.x) >= x &&
			aocmath.Min(line1.a.y, line1.b.y) <= y && aocmath.Max(line1.a.y, line1.b.y) >= y &&
			aocmath.Min(line2.a.x, line2.b.x) <= x && aocmath.Max(line2.a.x, line2.b.x) >= x &&
			aocmath.Min(line2.a.y, line2.b.y) <= y && aocmath.Max(line2.a.y, line2.b.y) >= y {
			return DO_INTERSECT, point{x, y}
		}

		return DONT_INTERSECT, point{x, y}
	}
}

func ParseInput(problem string) []line {
	lineStrings := strings.Split(problem, "\n")
	lines := make([]line, len(lineStrings))

	for i, lineStr := range lineStrings {
		numbers := aocstring.StringArrayToIntArray(expr.FindAllString(lineStr, -1))
		lines[i] = line{
			point{
				numbers[0],
				numbers[1],
			},
			point{
				numbers[2],
				numbers[3],
			},
		}
	}

	return lines
}

func SolvePart1Second(problem string) int {
	lines := ParseInput(problem)
	maxX, maxY := GetDimensionsMap(lines)

	mineMap := CreateMap(maxX, maxY)

	for _, l := range lines {
		applyLine(&mineMap, l)
	}

	return countIntersections(&mineMap)
}

func countIntersections(mineMap *[][]int) int {
	counter := 0

	for i := 0; i < len(*mineMap); i++ {
		for j := 0; j < len((*mineMap)[0]); j++ {
			if (*mineMap)[i][j] > 1 {
				counter++
			}
		}
	}

	return counter
}

func applyLine(mineMap *[][]int, l line) {
	dir := l.direction()

	switch dir {
	case HORIZONTAL:
		for i := l.minX(); i < l.maxX(); i++ {
			(*mineMap)[l.a.y][i]++
		}
		break
	case VERTICAL:
		for i := l.minY(); i < l.maxY(); i++ {
			(*mineMap)[i][l.a.x]++
		}
		break
	}
}

func CreateMap(maxX, maxY int) [][]int {
	mineMap := make([][]int, maxY)

	for i := 0; i < maxX; i++ {
		mineMap[i] = make([]int, maxX, maxX)
	}

	return mineMap
}

func GetDimensionsMap(lines []line) (maxX, maxY int) {

	maxX = math.MinInt
	maxY = math.MinInt

	for _, line := range lines {
		maxX = aocmath.Max(maxX, line.maxX())
		maxY = aocmath.Max(maxY, line.maxY())
	}

	maxX++
	maxY++

	return
}
