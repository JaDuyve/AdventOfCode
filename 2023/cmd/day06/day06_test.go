package main

import (
	"testing"
)

func TestDay06(t *testing.T) {
	const input string = "Time:      7  15   30\nDistance:  9  40  200"

	t.Run("part 1", func(t *testing.T) {
		const expected int = 288
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 0
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}
