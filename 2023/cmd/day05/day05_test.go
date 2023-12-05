package main

import (
	"testing"
)

func TestDay05(t *testing.T) {
	const input string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	t.Run("part 1", func(t *testing.T) {
		const expected int = 35
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 46
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}
