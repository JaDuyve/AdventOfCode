package day03

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day03/input")
	return solvePart1(input)
}

func solvePart1(puzzleInput string) int {
	area := slices.ToTwoDimArray(strings.Split(puzzleInput, "\n"))

	numberOfTrees := calcNumberOfTreesForSlope1(area, 3, 1)

	return numberOfTrees
}

func calcNumberOfTreesForSlope1(area [][]rune, right int, down int) (numberOfTrees int) {

	areaWidth := len(area[0])

	for x, y := right, down; y < len(area); x, y = (x+right)%areaWidth, y+down {
		if area[y][x] == '#' {
			numberOfTrees++
		}
	}

	return
}
