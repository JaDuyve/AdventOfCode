package main

import (
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"time"

	"aoc/internal/utils"
)

const (
	HIGH_CARD     = -1
	ONE_PAIR      = 0
	TWO_PAIR      = 1
	THREE_OF_KIND = 2
	FULL_HOUSE    = 3
	FOUR_OF_KIND  = 4
	FIVE_OF_KIND  = 5
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

	if score[0] > 0 {
		jcount := score[0]

		max := math.MinInt
		max2 := math.MinInt
		index := -1
		index2 := -1
		for i := 0; i < len(score); i++ {
			if max < score[i] {
				index2 = index
				index = i
				max = score[i]
			} else if max2 < score[i] {
				index2 = i
				max2 = score[i]
			}
		}

		if index != 0 {
			score[index] += jcount
		} else {
			score[index2] += jcount
		}

		score[0] = 0
	}

	sort.Ints(score)

	if score[13] == 5 {
		return FIVE_OF_KIND
	}

	if score[13] == 4 {
		return FOUR_OF_KIND
	}

	if score[13] == 3 && score[12] == 2 {
		return FULL_HOUSE
	}

	if score[13] == 3 && score[12] == 1 {
		return THREE_OF_KIND
	}

	if score[13] == 2 && score[12] == 2 {
		return TWO_PAIR
	}

	if score[13] == 2 && score[12] == 1 {
		return ONE_PAIR
	}

	return HIGH_CARD
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
		//log.Printf("Cards: %s, Bid: %d\n", hands[i].Cards, hands[i].Bid)

	}

	return sum
}
