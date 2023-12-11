package main

import (
	"aoc/internal/utils"
	"fmt"
	"testing"
)

func TestDay11(t *testing.T) {
	const input string = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	t.Run("part 1", func(t *testing.T) {
		const expected int = 374
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	tests := []struct {
		expectedValue int
		multi         int
	}{
		{374, 2},
		{1030, 10},
		{8410, 100},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Part2 - test %d", i), func(t *testing.T) {
			actual := part2(input, tt.multi)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}

func TestCalcManhattenDistance(t *testing.T) {
	tests := []struct {
		expectedValue int
		p1            utils.Tuple[int, int]
		p2            utils.Tuple[int, int]
	}{
		{9, utils.Tuple[int, int]{A: 1, B: 6}, utils.Tuple[int, int]{A: 5, B: 11}},
		{15, utils.Tuple[int, int]{A: 4, B: 0}, utils.Tuple[int, int]{A: 9, B: 10}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("ManhattenDistance TestCase %d", i), func(t *testing.T) {
			actual := utils.ManhattenDistance(tt.p1, tt.p2)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}
