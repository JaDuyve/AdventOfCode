package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput(t *testing.T) {
	input := "class: 1-3 or 5-7\n" +
		"row: 6-11 or 33-44\n" +
		"seat: 13-40 or 45-50\n" +
		"\n" +
		"your ticket:\n" +
		"7,1,14\n" +
		"\n" +
		"nearby tickets:\n" +
		"7,3,47\n" +
		"40,4,50\n" +
		"55,2,20\n" +
		"38,6,12"

	expected := 71
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInput(t *testing.T) {
	input := "departure class: 0-1 or 4-19\n" +
		"row: 0-5 or 8-19\n" +
		"departure seat: 0-13 or 16-19\n" +
		"\n" +
		"your ticket:\n" +
		"11,12,13\n" +
		"\n" +
		"nearby tickets:\n" +
		"3,9,18\n" +
		"15,1,5\n" +
		"5,14,9"
	expected := 156
	assert.Equal(t, expected, solvePart2(input))
}

func Test_FilterNearbyTicket(t *testing.T) {
	input := "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12"
	fields, _, nearbyTickets := parseInput(input)

	filterNearbyTickets(fields, &nearbyTickets)

	expected := [][]int{{7, 3, 47}}
	assert.Equal(t, expected, nearbyTickets)
}
