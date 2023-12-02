package day07

import (
	"AdventOfCode2021/util/aocmath"
	"AdventOfCode2021/util/aocstring"
	"log"
	"math"
)

const (
	inputFile = "calendar/day07/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 07 part 1: The Treachery of Whales : %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 07 part 2: The Treachery of Whales : %d", solvePart2(problem))
}

func solvePart1(problem string) int {
	crabs := aocstring.SplitToIntArray(problem, ",")

	min := aocmath.MinOfArray(crabs)
	max := aocmath.MaxOfArray(crabs)

	fuel := math.MaxInt

	for i := min; i <= max; i++ {
		tmp := 0

		for _, number := range crabs {
			tmp += aocmath.Abs(number - i)
		}

		fuel = aocmath.Min(fuel, tmp)
	}

	return fuel
}

func solvePart2(problem string) int {
	crabs := aocstring.SplitToIntArray(problem, ",")

	min := aocmath.MinOfArray(crabs)
	max := aocmath.MaxOfArray(crabs)

	fuel := math.MaxInt

	for i := min; i <= max; i++ {
		tmp := 0

		for _, number := range crabs {
			tmp += calcFuel(aocmath.Abs(number-i), 1)
		}

		fuel = aocmath.Min(fuel, tmp)
	}

	return fuel
}

func calcFuel(number int, i int) int {

	for number == 0 {
		return 0
	}

	return calcFuel(number-1, i+1) + i
}
