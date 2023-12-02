package day01

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part1() int {
	numbers := files.ReadFile("calendar/day01/input")
	return solvePart1(numbers)
}

func solvePart1(puzzleInput string) int {
	expenseReport := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))

	for indexA, expenseA := range expenseReport {
		for indexB, expenseB := range expenseReport {
			if indexA == indexB {
				continue
			}

			if expenseA+expenseB == 2020 {
				return expenseA * expenseB
			}
		}
	}

	return -1
}
