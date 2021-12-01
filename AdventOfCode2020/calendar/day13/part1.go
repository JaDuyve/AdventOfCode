package day13

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day13/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	input := strings.Split(puzzleInput, "\n")

	departTime := conv.ToInt(input[0])
	busIds := strings.Split(input[1], ",")

	bestBus := 0
	bestTime := departTime
	for _, busIdString := range busIds {
		if busIdString == "x" {
			continue
		}

		busId := conv.ToInt(busIdString)

		time := ((departTime/busId + 1) * busId) - departTime
		if time < bestTime {
			bestBus = busId
			bestTime = time
		}
	}

	return bestBus * bestTime
}
