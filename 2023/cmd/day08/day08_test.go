package main

import (
	"fmt"
	"testing"
)

func TestDay08(t *testing.T) {

	tests := []struct {
		expectedValue int
		input         string
	}{
		{2, "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)"},
		{6, "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"},
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
		const input string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
		const expected int = 6
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}
