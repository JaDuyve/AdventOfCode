package utils

import "math"

func Max(arr []int) int {
	max := math.MinInt
	for _, val := range arr {
		if max < val {
			max = val
		}
	}

	return max
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}
