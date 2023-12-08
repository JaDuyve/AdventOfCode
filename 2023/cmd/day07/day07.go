package main

import (
	"log"
	"os"
	"slices"
	"sort"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 07, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input)
	us := time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res1, us)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input)
	us = time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res2, us)
}

var CardOrder = map[uint8]int8{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var CardOrder2 = map[uint8]int8{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

// part one
func part1(input string) int {
	handStr := utils.SplitToStringGroups(input, "\n", " ")

	hands := make([]Hand, len(handStr))

	for i := 0; i < len(hands); i++ {
		hands[i] = Hand{
			Cards: handStr[i][0],
			Bid:   utils.ParseInt(handStr[i][1]),
			Value: CalcHandValue(handStr[i][0]),
		}
	}

	slices.SortFunc(hands, HandComp)

	sum := 0

	for i := 0; i < len(hands); i++ {
		sum += hands[i].Bid * (i + 1)
	}

	return sum
}

type Hand struct {
	Cards string
	Bid   int
	Value int
}

func CalcHandValue2(cards string) int {
	score := make([]int, 14)

	for i := 0; i < len(cards); i++ {
		scoreIndex, _ := CardOrder2[cards[i]]

		score[scoreIndex]++
	}

	jcount := score[0]

	sort.Ints(score)

	score[len(score)-1] += jcount

	if score[13] == 5 {
		return 5
	}

	if score[13] == 4 {
		return 4
	}

	if score[13] == 3 && score[12] == 2 {
		return 3
	}

	if score[13] == 3 && score[12] == 1 {
		return 2
	}

	if score[13] == 2 && score[12] == 2 {
		return 1
	}

	if score[13] == 2 && score[12] == 1 {
		return 0
	}

	return -1
}

func CalcHandValue(cards string) int {
	score := make([]int, 14)

	for i := 0; i < len(cards); i++ {
		scoreIndex, _ := CardOrder[cards[i]]

		score[scoreIndex]++
	}

	sort.Ints(score)

	if score[13] == 5 {
		return 5
	}

	if score[13] == 4 {
		return 4
	}

	if score[13] == 3 && score[12] == 2 {
		return 3
	}

	if score[13] == 3 && score[12] == 1 {
		return 2
	}

	if score[13] == 2 && score[12] == 2 {
		return 1
	}

	if score[13] == 2 && score[12] == 1 {
		return 0
	}

	return -1
}

func HandComp(a, b Hand) int {
	if a.Value < b.Value {
		return -1
	}

	if b.Value < a.Value {
		return 1
	}

	for i := 0; i < len(a.Cards); i++ {
		valueA, _ := CardOrder[a.Cards[i]]
		valueB, _ := CardOrder[b.Cards[i]]

		if valueA < valueB {
			return -1
		}

		if valueA > valueB {
			return 1
		}
	}

	return 0
}

func HandComp2(a, b Hand) int {
	if a.Value < b.Value {
		return -1
	}

	if b.Value < a.Value {
		return 1
	}

	for i := 0; i < len(a.Cards); i++ {
		valueA, _ := CardOrder2[a.Cards[i]]
		valueB, _ := CardOrder2[b.Cards[i]]

		if valueA < valueB {
			return -1
		}

		if valueA > valueB {
			return 1
		}
	}

	return 0
}

// part two
func part2(input string) int {
	handStr := utils.SplitToStringGroups(input, "\n", " ")

	hands := make([]Hand, len(handStr))

	for i := 0; i < len(hands); i++ {
		hands[i] = Hand{
			Cards: handStr[i][0],
			Bid:   utils.ParseInt(handStr[i][1]),
			Value: CalcHandValue2(handStr[i][0]),
		}
	}

	slices.SortFunc(hands, HandComp2)

	sum := 0

	for i := 0; i < len(hands); i++ {
		sum += hands[i].Bid * (i + 1)
	}

	return sum
}
