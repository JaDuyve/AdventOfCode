package day14

import (
	"AdventOfCode2021/util/aocstring"
	"log"
)

const (
	inputFile = "calendar/day14/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 14 part 1: Extended Polymerization: %d", solvePart1(problem))
}

func solvePart1(problem string) int {
	poly := New(problem, 10)

	poly.ApplySteps()

	return poly.GetMinMaxDifference()
}

func part2(problem string) {
	log.Printf("Day 14 part 2: Extended Polymerization: %d", solvePart2(problem))
}

func solvePart2(problem string) int {
	poly := New(problem, 40)

	poly.ApplySteps()

	return poly.GetMinMaxDifference()
}
