package utils

func Remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Any[T any](s []T, fn func(T) bool) bool {
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			return true
		}
	}

	return false
}
