package day10

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day10/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	adapters := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))
	adapters = makeAdapterChain(adapters)

	adapterMap := make(map[int][]int)

	for _, element := range adapters {
		adapterMap[element] = []int{element + 3, element + 2, element + 1}
	}

	return countNumberOfDistinctAdapterChains(adapterMap, make(map[int]int), aocmath.Max(adapters), 0)
}

func countNumberOfDistinctAdapterChains(adapters map[int][]int, history map[int]int, target int, currentPos int) int {
	if value, seen := history[currentPos]; seen {
		return value
	}

	value := 0
	for _, current := range adapters[currentPos] {
		if current != target {
			value += countNumberOfDistinctAdapterChains(adapters, history, target, current)
			continue
		}

		value++
	}

	history[currentPos] = value
	return value
}
