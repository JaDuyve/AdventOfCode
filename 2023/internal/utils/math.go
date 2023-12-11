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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattenDistance(p1, p2 Tuple[int, int]) int {
	return Abs(p2.A-p1.A) + Abs(p2.B-p1.B)
}
