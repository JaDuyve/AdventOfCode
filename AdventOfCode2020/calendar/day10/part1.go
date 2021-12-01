package day10

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"sort"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day10/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	adapters := conv.ToIntSlice(strings.Split(puzzleInput, "\n"))
	adapters = makeAdapterChain(adapters)

	joltFreq := calculateJoltDifference(adapters)

	return joltFreq[1] * joltFreq[3]
}

func calculateJoltDifference(adapterChain []int) []int {
	freq := make([]int, 4)

	for i := 0; i < len(adapterChain)-1; i++ {
		joltDifference := adapterChain[i+1] - adapterChain[i]
		freq[joltDifference]++
	}

	return freq
}

func makeAdapterChain(adapters []int) []int {
	max := aocmath.Max(adapters)
	adapters = append(adapters, 0)
	adapters = append(adapters, max+3)

	sort.Ints(adapters)
	return adapters
}
