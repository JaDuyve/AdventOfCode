package main

import (
	"fmt"
	"testing"
)

func TestDay10(t *testing.T) {

	tests := []struct {
		expectedValue int
		input         string
	}{
		{4, "-L|F7\n7S-7|\nL|7||\n-L-J|\nL|-JF"},
		{8, "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Part 1 test %d", i), func(t *testing.T) {
			actual := part1(tt.input)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}

	t.Run("part 2", func(t *testing.T) {
		const input string = ""
		const expected int = 0
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}

func TestCalcLoopCount(t *testing.T) {
	tests := []struct {
		name          string
		x             int
		y             int
		expectedValue int
		maze          []string
	}{
		{"small loop", 1, 1, 8, []string{".....", ".F-7.", ".|.|.", ".L-J.", "....."}},
		{"no loop 1", 0, 0, 0, []string{"7-F7-", ".FJ|7", "SJLL7", "|F--J", "LJ.LJ"}},
		{"no loop 2", 5, 5, 0, []string{"7-F7-", ".FJ|7", "SJLL7", "|F--J", "LJ.LJ"}},
		{"longer no loop 1", 5, 0, 0, []string{"-L|F7", "7S-7|", "L|7||", "-L-J|", "L|-JF"}},
		//{"longer no loop 2", 2, 0, 0, []string{"-L|F7", "7S-7|", "L|7||", "-L-J|", "L|-JF"}},
		{"complex loop", 0, 2, 16, []string{"..F7.", ".FJ|.", "FJ.L7", "|F--J", "LJ..."}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calcLoopCount(tt.x, tt.y, &tt.maze)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}
