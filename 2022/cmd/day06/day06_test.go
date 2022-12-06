package main

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {

	testcases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := part1(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %d , actual: %d", tc.expected, actual)
			}
		})
	}
}

func TestDay06Part2(t *testing.T) {

	testcases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := part2(tc.input)
			if actual != tc.expected {
				t.Errorf("expected: %d , actual: %d", tc.expected, actual)
			}
		})
	}
}
