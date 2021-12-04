package day03

import (
	"AdventOfCode2021/util/aocstring"
	"errors"
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

	gammaRate := CalculateGammaRate(diagnosticNumbers, int(len(diagnosticReport[0])))
	epsilonRate := CalculateEpsilonRate(diagnosticNumbers, int(len(diagnosticReport[0])))

	//log.Printf("gamma rate: '%d' (binary '%b')", gammaRate, gammaRate)
	//log.Printf("epsilon rate: '%d' (binary '%b')", epsilonRate, epsilonRate)

	return gammaRate * epsilonRate
}

func convertBinaryStrings(binaryStrings []string) []int {
	numbers := make([]int, len(binaryStrings))

	for i, binaryString := range binaryStrings {
		if number, err := strconv.ParseInt(binaryString, 2, 64); err != nil {
			log.Fatal(err)
		} else {
			numbers[i] = int(number)
		}
	}

	return numbers
}

func CalculateGammaRate(numbers []int, size int) int {
	return CalculateValue(numbers, size, isCounter1LargerCounter0)
}

func CalculateEpsilonRate(numbers []int, size int) int {
	return CalculateValue(numbers, size, isCounter0LargerCounter1)
}

func isCounter1LargerCounter0(counter0, counter1 int) int {
	if counter1 > counter0 {
		return 1
	}
	return 0
}

func isCounter0LargerCounter1(counter0, counter1 int) int {
	if counter1 < counter0 {
		return 1
	}

	return 0
}

func isCounter1LargerOrEqualCounter0(counter0, counter1 int) int {
	if counter1 > counter0 || counter0 == counter1 {
		return 1
	}
	return 0
}

func isCounter0LargerOrEqualCounter1(counter0, counter1 int) int {
	if counter1 > counter0 || counter0 == counter1 {
		return 0
	}

	return 1
}

func CalculateValue(numbers []int, size int, f func(int, int) int) int {
	endValue := 0

	for i := size - 1; i >= 0; i-- {
		endValue = endValue << 1

		counter0, counter1 := counter0And1(numbers, i)

		endValue += f(counter0, counter1)
	}

	return endValue
}

func calculateOxygenGeneratorRating(numbers []int, size int) int {
	return calculateComplexValue(numbers, size, isCounter1LargerOrEqualCounter0)
}

func calculateCo2ScrubberRating(numbers []int, size int) int {
	return calculateComplexValue(numbers, size, isCounter0LargerOrEqualCounter1)
}

func calculateComplexValue(numbers []int, size int, f func(int, int) int) int {
	if len(numbers) == 1 {
		return int(numbers[0])
	} else if size < 0 {
		log.Fatalf("no metric number found")
	}

	filteredNumbers, err := filterNumbers(numbers, size, f)
	if err != nil {
		log.Fatal(err)
	}

	return calculateComplexValue(filteredNumbers, size-1, f)
}

func filterNumbers(numbers []int, bitPos int, f func(int, int) int) ([]int, error) {
	m := int(1 << (bitPos - 1))

	numbers0 := make([]int, 0)
	numbers1 := make([]int, 0)

	for _, number := range numbers {
		r := number & m

		if r == m {
			numbers1 = append(numbers1, number)
		} else {
			numbers0 = append(numbers0, number)
		}
	}

	switch f(len(numbers0), len(numbers1)) {
	case 0:
		return numbers0, nil
	case 1:
		return numbers1, nil
	default:
		return nil, errors.New("can't filter number array")
	}
}

func counter0And1(numbers []int, bitPos int) (counter0, counter1 int) {
	m := int(1 << bitPos)

	for _, number := range numbers {
		r := number & m

		if r == m {
			counter1++
		} else {
			counter0++
		}
	}

	return
}

func part2(problem string) {
	log.Printf("Day 03 part 2: Binary Diagnostic: %d", solvePart2(problem))
}

func solvePart2(diagnosticReportString string) int {
	diagnosticReport := strings.Split(diagnosticReportString, "\n")

	diagnosticNumbers := convertBinaryStrings(diagnosticReport)

	oxygenGeneratorRating := calculateOxygenGeneratorRating(diagnosticNumbers, int(len(diagnosticReport[0])))
	co2ScrubberRating := calculateCo2ScrubberRating(diagnosticNumbers, int(len(diagnosticReport[0])))

	//log.Printf("gamma rate: '%d' (binary '%b')", oxygenGeneratorRating, oxygenGeneratorRating)
	//log.Printf("epsilon rate: '%d' (binary '%b')", co2ScrubberRating, co2ScrubberRating)

	return oxygenGeneratorRating * co2ScrubberRating
}
