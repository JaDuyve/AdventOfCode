package day03

import (
	"AdventOfCode2021/util/aocstring"
	"log"
	"strconv"
	"strings"
)

const (
	inputFile = "calendar/day03/input"
)

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 03 part 1: Binary Diagnostic 'Power Consumption': %d", solvePart1(problem))
}

func solvePart1(diagnosticReportString string) int {
	diagnosticReport := strings.Split(diagnosticReportString, "\n")

	diagnosticNumbers := convertBinaryStrings(diagnosticReport)

	gammaRate := CalculateGammaRate(diagnosticNumbers, int64(len(diagnosticReport[0])))
	epsilonRate := CalculateEpsilonRate(diagnosticNumbers, int64(len(diagnosticReport[0])))

	//log.Printf("gamma rate: '%d' (binary '%b')", gammaRate, gammaRate)
	//log.Printf("epsilon rate: '%d' (binary '%b')", epsilonRate, epsilonRate)

	return gammaRate * epsilonRate
}

func convertBinaryStrings(binaryStrings []string) []int64 {
	numbers := make([]int64, len(binaryStrings))

	for i, binaryString := range binaryStrings {
		if number, err := strconv.ParseInt(binaryString, 2, 64); err != nil {
			log.Fatal(err)
		} else {
			numbers[i] = number
		}
	}

	return numbers
}

func CalculateGammaRate(numbers []int64, size int64) int {
	gammaRateFunc := func(counter1, counter0 int) bool { return counter1 > counter0 }

	return CalculateValue(numbers, size, gammaRateFunc)
}

func CalculateEpsilonRate(numbers []int64, size int64) int {
	epsilonRateFunc := func(counter1, counter0 int) bool { return counter1 < counter0 }

	return CalculateValue(numbers, size, epsilonRateFunc)
}

func CalculateValue(numbers []int64, size int64, f func(int, int) bool) int {
	endValue := 0

	for i := size - 1; i >= 0; i-- {
		endValue = endValue << 1

		counter1 := 0
		counter0 := 0
		m := int64(1 << i)

		for _, number := range numbers {
			r := number & m

			if r == m {
				counter1++
			} else {
				counter0++
			}
		}

		if f(counter1, counter0) {
			endValue++
		}
	}

	return endValue
}

func part2(problem string) {
	log.Printf("Day 03 part 2: Binary Diagnostic: %d", solvePart2(problem))
}

func solvePart2(diagnosticString string) int {
	return 0
}
