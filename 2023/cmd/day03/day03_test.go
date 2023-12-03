package main

import (
	"testing"
)

func TestDay03(t *testing.T) {
	const input string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	t.Run("part 1", func(t *testing.T) {
		const expected int = 925 //413 //4361
		actual := part1(input)
		t.Log(actual)

		if actual != expected {
			t.Errorf("expected: %d , actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 0
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d , actual: %d", expected, actual)
		}
	})
}
