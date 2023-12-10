package main

import (
	"testing"
)

func TestDay09(t *testing.T) {
	const input string = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"

	t.Run("part 1", func(t *testing.T) {
		const expected int = 114
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 2
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}

func TestCalcHistory(t *testing.T) {
	tests := []struct {
		test          string
		expectedValue int
		history       []int
	}{
		{"0 3 6 9 12 15", 18, []int{0, 3, 6, 9, 12, 15}},
		{"1 3 6 10 15 21", 28, []int{1, 3, 6, 10, 15, 21}},
		{"10 13 16 21 30 45", 68, []int{10, 13, 16, 21, 30, 45}},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			actual := calcHistoryValue(tt.history)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}

func TestCalcHistory2(t *testing.T) {
	tests := []struct {
		test          string
		expectedValue int
		history       []int
	}{
		{"0 3 6 9 12 15", -3, []int{0, 3, 6, 9, 12, 15}},
		{"1 3 6 10 15 21", 0, []int{1, 3, 6, 10, 15, 21}},
		{"10 13 16 21 30 45", 5, []int{10, 13, 16, 21, 30, 45}},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			actual := calcHistoryValue2(tt.history)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}
