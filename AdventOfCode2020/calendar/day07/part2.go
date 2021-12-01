package day07

import (
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day07/input")
	return solvePart2(input)
}

func solvePart2(puzzleInput string) int {
	rules := strings.Split(puzzleInput, "\n")
	bags := buildBags(rules)

	return countNeededBagsFor("shiny gold", &bags)
}

func countNeededBagsFor(color string, bags *map[string]Bag) int {
	sum := 0

	for _, inner := range (*bags)[color].contains {
		sum += inner.amount + inner.amount*countNeededBagsFor(inner.color, bags)

	}

	return sum
}
