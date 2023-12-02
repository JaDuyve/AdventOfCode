package day04

import (
	"AdventOfCode2020/utils/aocstring"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"regexp"
	"strings"
)

var eyeColorRegex = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var hairColorRegex = regexp.MustCompile("^#[0-9a-f]{6}$")
var passportIdRegex = regexp.MustCompile("^[0-9]{9}$")

var fields = map[string]func(string) bool{
	"byr": isValidBirthYear,
	"iyr": isValidIssueYear,
	"eyr": isValidExpirationYear,
	"hgt": isValidHeight,
	"hcl": isValidHairColor,
	"ecl": isValidEyeColor,
	"pid": isValidPassportId,
}

func Part2() int {
	input := files.ReadFile("calendar/day04/input")

	return solvePart2(input)
}

func solvePart2(puzzleInput string) int {
	passports := strings.Split(puzzleInput, "\n\n")
	return len(slices.StringFilter(passports, isPassportValidWithConstraints))
}

func isPassportValidWithConstraints(passportString string) bool {
	passportString = strings.ReplaceAll(passportString, "\n", " ")
	passport := aocstring.ExtractFieldOfString(strings.Split(passportString, " "), ":")

	for field, validationFunction := range fields {
		value, ok := passport[field]

		if ok == false {
			return false
		}

		if !validationFunction(value) {
			return false
		}
	}

	return true
}

func isValidBirthYear(year string) bool {
	return isValidYear(year, 1920, 2002)
}

func isValidIssueYear(year string) bool {
	return isValidYear(year, 2010, 2020)
}

func isValidExpirationYear(year string) bool {
	return isValidYear(year, 2020, 2030)
}

func isValidYear(yearString string, start int, end int) bool {
	year := conv.ToInt(yearString)
	return year >= start && year <= end
}

func isValidEyeColor(color string) bool {
	return eyeColorRegex.MatchString(color)
}

func isValidHairColor(color string) bool {
	return hairColorRegex.MatchString(color)
}

func isValidPassportId(str string) bool {
	return passportIdRegex.MatchString(str)
}

func isValidHeight(str string) bool {

	switch {
	case str[len(str)-2:] == "cm":
		height := conv.ToInt(str[:len(str)-2])

		if height >= 150 && height <= 193 {
			return true
		}
	case str[len(str)-2:] == "in":
		height := conv.ToInt(str[:len(str)-2])

		if height >= 59 && height <= 76 {
			return true
		}
	default:
		return false
	}

	return false
}
