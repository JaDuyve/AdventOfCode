package day01

import (
	"AdventOfCode2017/utils/files"
	"log"
)

func Run() {
	part1()
	part2()
}

func part1() {
	inputString := files.ReadFile("calendar/day01/input1")

	solution := solveProblem1(inputString)

	log.Printf("Day 01 part 2 solution: %d", solution)
}

func solveProblem1(input string) int {
	var sum int

	firstNumber := getValue(0, input)
	for i := 1; i <= len(input); i++ {
		secondNumber := getValue(i, input)

		if firstNumber == secondNumber {
			sum += firstNumber
		}

		firstNumber = secondNumber
	}

	return sum
}

func part2() {
	inputString := files.ReadFile("calendar/day01/input1")

	solution := solveProblem2(inputString)

	log.Printf("Day 01 part 1 solution: %d", solution)
}

func solveProblem2(input string) int {
	var sum int

	halfWayLength := len(input) / 2

	for i := 0; i < len(input); i++ {
		firstNumber := getValue(i, input)
		secondNumber := getValue(i+halfWayLength, input)

		if firstNumber == secondNumber {
			sum += firstNumber
		}
	}

	return sum
}

func getValue(index int, input string) int {

	return int(input[index%len(input)] - '0')
}
