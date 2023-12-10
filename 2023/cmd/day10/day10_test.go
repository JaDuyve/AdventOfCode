package main

import (
	"fmt"
	"testing"
)

func TestDay10(t *testing.T) {

	testsPart1 := []struct {
		expectedValue int
		input         string
	}{
		{4, "-L|F7\n7S-7|\nL|7||\n-L-J|\nL|-JF"},
		{8, "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"},
	}

	for i, tt := range testsPart1 {
		t.Run(fmt.Sprintf("Part 1 test %d", i), func(t *testing.T) {
			actual := part1(tt.input)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}

	testsPart2 := []struct {
		expectedValue int
		input         string
	}{
		{4, "...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n..........."},
		{4, "..........\n.S------7.\n.|F----7|.\n.||....||.\n.||....||.\n.|L-7F-J|.\n.|..||..|.\n.L--JL--J.\n.........."},
		{8, ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ..."},
		{10, "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L"},
	}

	for i, tt := range testsPart2 {
		t.Run(fmt.Sprintf("Part 2 test %d", i), func(t *testing.T) {
			actual := part2(tt.input)

			if actual != tt.expectedValue {
				t.Errorf("expected: %d, actual: %d", tt.expectedValue, actual)
			}
		})
	}
}
