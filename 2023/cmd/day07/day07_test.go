package main

import (
	"testing"
)

func TestDay07(t *testing.T) {
	const input string = "2345A 1\nQ2KJJ 13\nQ2Q2Q 19\nT3T3J 17\nT3Q33 11\n2345J 3\nJ345A 2\n32T3K 5\nT55J5 29\nKK677 7\nKTJJT 34\nQQQJA 31\nJJJJJ 37\nJAAAA 43\nAAAAJ 59\nAAAAA 61\n2AAAA 23\n2JJJJ 53\nJJJJ2 41"

	t.Run("part 1", func(t *testing.T) {
		const expected int = 6592
		actual := part1(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		const expected int = 6839
		actual := part2(input)

		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}

func TestCalcHandValue2(t *testing.T) {
	tests := []struct {
		cards         string
		expectedScore int
	}{
		{"2233J", FULL_HOUSE},
		{"2JJJJ", FIVE_OF_KIND},
		{"J345A", ONE_PAIR},

		{"J2345", ONE_PAIR},
		{"32T3K", ONE_PAIR},
		{"KK677", TWO_PAIR},
		{"T55J5", FOUR_OF_KIND},
		{"KTJJT", FOUR_OF_KIND},
		{"QQQJA", FOUR_OF_KIND},
		{"JJJJJ", FIVE_OF_KIND},
		{"96J66", FOUR_OF_KIND},
		{"J6569", THREE_OF_KIND},
		{"JKK92", THREE_OF_KIND},
	}

	for _, tt := range tests {
		t.Run(tt.cards, func(t *testing.T) {
			actual := CalcHandValue2(tt.cards)

			if actual != tt.expectedScore {
				t.Errorf("expected: %d, actual: %d", tt.expectedScore, actual)
			}
		})
	}
}
