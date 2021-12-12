package day09

import (
	"AdventOfCode2021/util/aocstring"
	"AdventOfCode2021/util/queue"
	"log"
	"strings"
)

const (
	inputFile = "calendar/day09/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 09 part 1: Smoke Basin: %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 09 part 2: Smoke Basin: %d", solvePart2(problem))
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

func calculateSizeBasin(y, x int, heightMap *[][]int) int {
	walkedPath := make([][]bool, len(*heightMap))
	for y := 0; y < len(*heightMap); y++ {
		walkedPath[y] = make([]bool, len((*heightMap)[y]))
	}

	return calculateSizeBasinRec(y, x, heightMap, walkedPath)
}

func solvePart2(problem string) int {
	heightMap := parseInput(problem)
	q := queue.NewQueue(50)

	for y := range heightMap {
		for x := range heightMap[y] {
			if !isLowPoint(x, y, &heightMap) {
				continue
			}

			size := calculateSizeBasin(y, x, &heightMap)
			q.Add(size)
		}
	}

	if q.Size() < 3 {
		log.Fatalf("Not enough low points found, number of lowpoints %d", q.Size())
	}

	return q.Remove() * q.Remove() * q.Remove()
}

func calculateSizeBasinRec(y, x int, heightMap *[][]int, walkedPath [][]bool) int {
	if walkedPath[y][x] {
		return 0
	}

	walkedPath[y][x] = true

	if (*heightMap)[y][x] == 9 {
		return 0
	}

	counter := 0

	if x != 0 {
		counter += calculateSizeBasinRec(y, x-1, heightMap, walkedPath)
	}

	if x != len((*heightMap)[y])-1 {
		counter += calculateSizeBasinRec(y, x+1, heightMap, walkedPath)
	}

	if y != 0 {
		counter += calculateSizeBasinRec(y-1, x, heightMap, walkedPath)
	}

	if y != len(*heightMap)-1 {
		counter += calculateSizeBasinRec(y+1, x, heightMap, walkedPath)
	}

	return counter + 1
}
