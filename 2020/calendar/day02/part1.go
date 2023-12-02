package day02

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day02/input")
	return solvePart1(input)
}

func solvePart1(puzzleInput string) int {
	validPasswords := slices.StringFilter(strings.Split(puzzleInput, "\n"), isPasswordValidCountChar)
	return len(validPasswords)
}

func isPasswordValidCountChar(entry string) bool {
	var min, max int
	var character, password string

	_, err := fmt.Sscanf(entry, "%d-%d %1s: %s", &min, &max, &character, &password)
	if err != nil {
		panic(err)
	}

	count := strings.Count(password, character)
	return count >= min && count <= max
}
