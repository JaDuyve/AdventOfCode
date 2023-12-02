package aocmath

import "math"

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func MinOfArray(numbers []int) int {
	min := math.MaxInt

	for _, number := range numbers {
		min = Min(min, number)
	}

	return min
}

func MaxOfArray(numbers []int) int {
	max := math.MinInt

	for _, number := range numbers {
		max = Max(max, number)
	}

	return max
}
