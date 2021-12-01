package day06

import (
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day06/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	groups := strings.Split(puzzleInput, "\n\n")
	sum := 0

	for _, group := range groups {
		sum += countQuestions(group)
	}

	return sum
}

func countQuestions(questions string) int {
	numberOfPersons := strings.Count(questions, "\n") + 1
	q := strings.ReplaceAll(questions, "\n", "")
	uniqueQuestions := make(map[rune]int)

	for _, char := range q {
		uniqueQuestions[char]++
	}

	numberOfQuestions := 0
	for _, amount := range uniqueQuestions {
		if numberOfPersons == amount {
			numberOfQuestions++
		}
	}

	return numberOfQuestions
}
