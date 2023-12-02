package day15

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day15/input")
	return solvePart1(input, 2020)

}

func solvePart1(puzzleInput string, turn int) int {
	numbers := conv.ToIntSlice(strings.Split(puzzleInput, ","))
	initialSize := len(numbers)
	lastOccurrenceMap := make(map[int]int)

	for i := 0; i < turn; i++ {
		if i < initialSize {
			if i > 0 {
				lastOccurrenceMap[numbers[i-1]] = i - 1
			}
			continue
		}

		if previousOccurrence, exist := lastOccurrenceMap[numbers[i-1]]; !exist {
			numbers = append(numbers, 0)
		} else {
			numbers = append(numbers, (i-1)-previousOccurrence)
		}
		lastOccurrenceMap[numbers[i-1]] = i - 1
	}

	return numbers[turn-1]
}
