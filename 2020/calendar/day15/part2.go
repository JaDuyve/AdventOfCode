package day15

import (
	"AdventOfCode2020/utils/files"
)

func Part2() int {
	input := files.ReadFile("calendar/day15/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	return solvePart1(puzzleInput, 30_000_000)
}
