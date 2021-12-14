package day14

import (
	"AdventOfCode2021/util/aocstring"
	"log"
)

const (
	inputFile = "calendar/day11/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
}

func part1(problem string) {
	log.Printf("Day 11 part 1: Extended Polymerization: %d", solvePart1(problem))
}

func solvePart1(problem string) int {
	return 1
}
