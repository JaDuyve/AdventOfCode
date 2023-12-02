package day06

import (
	"AdventOfCode2021/util/aocstring"
	"log"
)

const (
	inputFile = "calendar/day06/input"

	cycle = 9
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 06 part 1: Lanternfish : %d", solvePart1(problem, 80))
}

func solvePart1(problem string, days int) int {
	freq := loadFishFreq(problem)

	for i := 0; i < days; i++ {
		applyDay(&freq)
	}

	sum := 0
	for _, n := range freq {
		sum += n
	}

	return sum
}

func applyDay(freq *[]int) {
	tmp := (*freq)[cycle-1]
	for i := cycle - 1; i > 0; i-- {
		t := (*freq)[i-1]
		(*freq)[i-1] = tmp
		tmp = t
	}

	(*freq)[6] += tmp
	(*freq)[8] = tmp
}

func loadFishFreq(lanternFishStr string) []int {
	lanternFish := aocstring.SplitToIntArray(lanternFishStr, ",")

	freq := make([]int, cycle)

	for _, fish := range lanternFish {
		freq[fish]++
	}

	return freq
}

func part2(problem string) {
	log.Printf("Day 06 part 2: Lanternfish : %d", solvePart1(problem, 256))
}
