package utils

import (
	"log"
	"strconv"
	"strings"
)

type Number interface {
	int | int64 | float64
}

func SplitToInt(string, sep string) []int {
	stringArr := strings.Split(string, sep)

	numberArr := make([]int, len(stringArr))
	for i, val := range stringArr {
		number, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Unable to parse string '%s' to int", val)
		}
		numberArr[i] = number
	}
	return numberArr
}

func SplitToIntGroups(data, sep, secondSep string) [][]int {
	splitInput := strings.Split(data, sep)

	groups := make([][]int, len(splitInput))
	for i, elf := range splitInput {
		groups[i] = SplitToInt(elf, secondSep)
	}

	return groups
}
