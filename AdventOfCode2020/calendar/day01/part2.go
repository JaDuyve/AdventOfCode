package day01

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	numbers := files.ReadFile("calendar/day01/input")
	return solvePart2(numbers)
}

func solvePart2(puzzleInput string) int {
	expenseReport := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))

	sum := aocmath.ThreeSum(expenseReport, 2020)[0]

	return sum[0] * sum[1] * sum[2]
}
