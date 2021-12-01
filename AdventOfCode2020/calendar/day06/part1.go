package day06

import (
	"AdventOfCode2020/utils/files"
	"go/types"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day06/input")
	return solvePart1(input)
}

func solvePart1(puzzleInput string) int {
	groups := strings.Split(puzzleInput, "\n\n")
	sum := 0

	for _, group := range groups {
		sum += countUniqueQuestions(group)
	}

	return sum
}

func countUniqueQuestions(questions string) int {
	q := strings.ReplaceAll(questions, "\n", "")
	uniqueQuestions := make(map[rune]types.Nil)

	for _, char := range q {
		uniqueQuestions[char] = types.Nil{}
	}

	return len(uniqueQuestions)
}
