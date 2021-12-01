package day02

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day02/input")
	return solvePart2(input)
}

func solvePart2(puzzleInput string) int {
	validPasswords := slices.StringFilter(strings.Split(puzzleInput, "\n"), isPasswordValid)

	return len(validPasswords)
}

func isPasswordValid(entry string) bool {
	var min, max int
	var character uint8
	var password string

	_, err := fmt.Sscanf(entry, "%d-%d %c: %s", &min, &max, &character, &password)
	if err != nil {
		panic(err)
	}

	return (password[min-1] == character) != (password[max-1] == character)
}
