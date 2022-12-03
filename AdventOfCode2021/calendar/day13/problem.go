package day13

import (
	"AdventOfCode2021/util/aocstring"
	"log"
)

const inputFile = "calendar/day13/input"

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 13 part 1: Transparent Origami: %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 13 part 2: Transparent Origami: \n%s", solvePart2(problem))
}

func solvePart1(problem string) int {
	p := NewPaper(problem)

	p.Fold()

	return p.CountDots()
}

func solvePart2(problem string) string {
	p := NewPaper(problem)

	for range p.Instructions {
		p.Fold()
	}

	return p.ToString()
}
