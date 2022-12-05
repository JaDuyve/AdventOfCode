package utils

type Number interface {
	int | int8 | int32 | int64 | float32 | float64
}

func MaxArr[T Number](arr []T) T {
	max := arr[0]
	for _, val := range arr[1:] {
		max = Max(max, val)
	}

	return max
}

func MinArr[T Number](arr []T) T {
	min := arr[0]
	for _, val := range arr[1:] {
		min = Min(min, val)
	}

	return min
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}
