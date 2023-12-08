package main

import (
	"testing"
)

func TestDay07(t *testing.T) {
	const input string = "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	t.Run("part 1", func(t *testing.T) {
		const expected int = 6440
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 5905
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}
