package day13

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day13/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	splitString := strings.Split(strings.Split(puzzleInput, "\n")[1], ",")

	numbers := extractNumbers(splitString)
	remainders := calculateRemainders(splitString)
	product := calculateProduct(numbers)
	pp := calculatePP(product, numbers)
	inverse := calculateInverse(pp, numbers)

	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += remainders[i] * pp[i] * inverse[i]
	}

	return sum % product
}

func extractNumbers(input []string) []int {
	numbers := make([]int, 0)

	for _, number := range input {
		if number == "x" {
			continue
		}

		numbers = append(numbers, conv.ToInt(number))
	}

	return numbers

}

func calculateRemainders(input []string) []int {
	remainders := make([]int, 0)
	remainders = append(remainders, 0)

	for i := 1; i < len(input); i++ {
		if input[i] == "x" {
			continue
		}

		number := conv.ToInt(input[i])
		remainders = append(remainders, number-i)
	}

	return remainders
}

func calculateProduct(numbers []int) int {
	product := numbers[0]

	for i := 1; i < len(numbers); i++ {
		product *= numbers[i]
	}

	return product
}

func calculateInverse(pp []int, numbers []int) []int {
	inverse := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {

		inverse[i] = aocmath.ModInverse(pp[i], numbers[i])
	}

	return inverse
}

func calculatePP(product int, numbers []int) []int {
	pp := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		pp[i] = product / numbers[i]
	}

	return pp
}
