package main

import (
	"testing"
)

func TestDayXX(t *testing.T) {
	input := ""

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d , actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d , actual: %d", expected, actual)
		}
	})
}
