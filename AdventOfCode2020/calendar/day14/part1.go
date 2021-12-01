package day14

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"regexp"
	"strconv"
	"strings"
)

var indexValueRegex = regexp.MustCompile("^mem\\W(\\d+)\\W = (\\d+)$")

func Part1() int {
	input := files.ReadFile("calendar/day14/input")
	return solvePart1(input)
}

func solvePart1(puzzleInput string) int {
	initializeProgram := strings.Split(puzzleInput, "\n")
	memory := make(map[int]int)
	var ones, zeros int64
	maskRegex := regexp.MustCompile("^mask = (\\w{36})$")

	for _, operation := range initializeProgram {
		if maskRegex.MatchString(operation) {
			mask := maskRegex.FindStringSubmatch(operation)[1]
			ones, _ = strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, 64)
			zeros, _ = strconv.ParseInt(strings.ReplaceAll(mask, "X", "1"), 2, 64)
			continue
		}

		index, value := getIndexValuePair(operation)
		memory[index] = (value | int(ones)) & int(zeros)

	}

	return calculateSum(memory)
}

func getIndexValuePair(operation string) (index int, value int) {
	values := indexValueRegex.FindAllStringSubmatch(operation, -1)[0]
	index = conv.ToInt(values[1])
	value = conv.ToInt(values[2])
	return
}

func calculateSum(memory map[int]int) int {
	sum := 0

	for _, value := range memory {
		sum += value
	}

	return sum
}
