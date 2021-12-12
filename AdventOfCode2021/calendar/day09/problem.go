package day09

import (
	"AdventOfCode2021/util/aocstring"
	"log"
	"strings"
)

const (
	inputFile = "calendar/day09/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
}

func part1(problem string) {
	log.Printf("Day 09 part 1: Smoke Basin: %d", solvePart1(problem))
}

func solvePart1(problem string) int {
	heightMap := parseInput(problem)

	lowPointsCounter := 0

	for y := range heightMap {
		for x := range heightMap[y] {
			if isLowPoint(x, y, &heightMap) {
				lowPointsCounter += 1 + heightMap[y][x]
			}
		}
	}

	return lowPointsCounter
}

func parseInput(problem string) [][]int {
	lines := strings.Split(problem, "\n")
	heightMap := make([][]int, len(lines))

	for i, l := range lines {
		heightMap[i] = aocstring.SplitToIntArray(l, "")
	}

	return heightMap
}

func isLowPoint(x, y int, heightMap *[][]int) bool {
	currentPoint := (*heightMap)[y][x]

	if x != 0 && currentPoint >= (*heightMap)[y][x-1] {
		return false
	}

	if x != len((*heightMap)[y])-1 && currentPoint >= (*heightMap)[y][x+1] {
		return false
	}

	if y != 0 && currentPoint >= (*heightMap)[y-1][x] {
		return false
	}

	if y != len(*heightMap)-1 && currentPoint >= (*heightMap)[y+1][x] {
		return false
	}

	return true
}
