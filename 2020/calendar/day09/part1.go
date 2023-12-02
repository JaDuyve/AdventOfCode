package day09

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day09/input")
	return solvePart1(input, 25)
}

func solvePart1(puzzleInput string, preamble int) int {
	cypher := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))

	return findInvalidNumber(cypher, preamble)
}

func findInvalidNumber(cypher []int, preamble int) int {
	for i := preamble; i < len(cypher); i++ {
		if !foundSum(cypher[i], cypher[i-preamble:i]) {
			return cypher[i]
		}
	}

	return -1
}

func foundSum(sum int, preambleRange []int) bool {
	for i := 0; i < len(preambleRange); i++ {
		for j := 0; j < len(preambleRange); j++ {
			if j == i {
				continue
			}

			if preambleRange[i]+preambleRange[j] == sum {
				return true
			}
		}
	}

	return false
}
