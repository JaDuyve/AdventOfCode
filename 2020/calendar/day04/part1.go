package day04

import (
	"AdventOfCode2020/utils/aocstring"
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day04/input")
	return solvePart1(input)
}

func solvePart1(puzzleInput string) int {
	passports := strings.Split(puzzleInput, "\n\n")
	return len(slices.StringFilter(passports, isPassportValid))
}

func isPassportValid(passportString string) bool {
	passportString = strings.ReplaceAll(passportString, "\n", " ")
	passport := aocstring.ExtractFieldOfString(strings.Split(passportString, " "), ":")

	for field, _ := range fields {
		if _, ok := passport[field]; ok == false {
			return false
		}
	}

	return true
}
