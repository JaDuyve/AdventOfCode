package utils

import (
	"log"
	"strconv"
	"strings"
)

func SplitTo[T any](data, sep string, parseFunc func(string) T) []T {
	group := strings.Split(data, sep)

	numberArr := make([]T, len(group))
	for i, item := range group {
		numberArr[i] = parseFunc(item)
	}
	return numberArr
}

func SplitToIgnoreSpace[T any](data, sep string, parseFunc func(string) T) []T {
	group := strings.Split(data, sep)

	var numberArr []T
	for _, item := range group {
		trimmedItem := strings.TrimSpace(item)
		if trimmedItem == "" {
			continue
		}

		numberArr = append(numberArr, parseFunc(trimmedItem))
	}
	return numberArr
}

func SplitToIntGroups(data, sep, secondSep string) [][]int {
	return SplitToGroups(data, sep, secondSep, ParseInt)
}

func SplitToStringGroups(data, sep, secondSep string) [][]string {
	return SplitToGroups(data, sep, secondSep, ParseString)
}

func SplitToGroups[T any](data, sep, secondSep string, parseFunc func(string) T) [][]T {
	splitInput := strings.Split(data, sep)

	groups := make([][]T, len(splitInput))
	for i, val := range splitInput {
		groups[i] = SplitTo(val, secondSep, parseFunc)
	}

	return groups
}

func ParseInt(val string) int {
	number, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Unable to parse string '%s' to int", val)
	}
	return number
}

func ParseString(val string) string {
	return val
}
