package day11

import (
	"AdventOfCode2021/calendar/day11/grid"
	"AdventOfCode2021/util/aocstring"
	"log"
)

const (
	inputFile = "calendar/day11/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 11 part 1: Dumbo Octopus: %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 11 part 2: Dumbo Octopus: %d", solvePart2(problem))
}

func solvePart1(problem string) int {
	g := grid.New(problem)

	return applySteps(g, 100)
}

func applySteps(g *grid.Grid, n int) int {
	flashes := 0

	for i := 0; i < n; i++ {
		flashes += g.ApplyStep()
	}

	return flashes
}

func solvePart2(problem string) int {
	step := 0

	g := grid.New(problem)

	for !g.IsSynchronised() {
		step++

		g.ApplyStep()
	}

	return step
}
