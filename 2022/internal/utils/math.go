package utils

type Number interface {
	int | int8 | int32 | int64 | float32 | float64
}

func Max[T Number](arr []T) T {
	max := arr[0]
	for _, val := range arr[1:] {
		if max < val {
			max = val
		}
	}

	return max
}

func Min[T Number](arr []T) T {
	min := arr[0]
	for _, val := range arr[1:] {
		if min > val {
			min = val
		}
	}

	return min
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}
