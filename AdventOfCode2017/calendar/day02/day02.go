package day02

import (
	"AdventOfCode2017/utils/files"
	"log"
	"math"
	"strconv"
	"strings"
)

func Run() {
	part1()
	part2()
}

func part1() {
	inputString := files.ReadFile("calendar/day02/input1")

	solution := solveProblem1(inputString)

	log.Printf("Day 02 part 1 solution: %d", solution)
}

func part2() {
	inputString := files.ReadFile("calendar/day02/input1")

	solution := solveProblem2(inputString)

	log.Printf("Day 02 part 2 solution: %d", solution)
}

func solveProblem1(input string) int {
	var sum int

	for _, row := range strings.Split(input, "\n") {
		min, max := minMax(row)

		sum += max - min
	}

	return sum
}

func solveProblem2(input string) int {
	var sum int

	for _, row := range strings.Split(input, "\n") {
		sum += getEvenlyDivisibleValue(row)
	}

	return sum
}

func minMax(input string) (int, int) {

	min := math.MaxInt32
	max := math.MinInt32

	for _, numberString := range strings.Split(input, "\t") {
		number, err := strconv.Atoi(numberString)

		if err != nil {
			log.Fatal("Failed to convert to number")
		}

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	return min, max
}

func getEvenlyDivisibleValue(input string) int {
	numberStrings := strings.Split(input, "\t")

	for indexA, numberStringA := range numberStrings {
		numberA, _ := strconv.Atoi(numberStringA)

		for indexB, numberStringB := range numberStrings {
			numberB, _ := strconv.Atoi(numberStringB)

			if numberA%numberB == 0 && indexA != indexB {
				return numberA / numberB
			}
		}
	}

	log.Fatalf("No evenly divisible value found for %s", input)
	return -1
}
