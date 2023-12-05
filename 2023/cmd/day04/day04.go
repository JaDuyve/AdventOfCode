package main

import (
	"log"
	"math"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 04, session)
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

type Card struct {
	WinNum  map[int]struct{}
	Numbers []int
}

func (card Card) CalcWins() int {
	found := 0
	for _, number := range card.Numbers {
		if _, ok := card.WinNum[number]; ok {
			found++
		}
	}

	if found == 1 {
		return 1
	}

	return int(math.Pow(2, float64(found-1)))
}

func (card Card) GetNumberOfWinningNumbers() int {
	found := 0
	for _, number := range card.Numbers {
		if _, ok := card.WinNum[number]; ok {
			found++
		}
	}

	return found
}

// part one
func part1(input string) int {
	cards := utils.SplitTo(input, "\n", parseCards)

	sum := 0
	for _, card := range cards {
		sum += card.CalcWins()
	}

	return sum
}

func parseCards(line string) Card {

	cardLine := utils.SplitToStringGroups(strings.Split(line, ":")[1], "|", " ")

	winNumbers := make(map[int]struct{})
	for _, number := range cardLine[0] {
		if number == "" {
			continue
		}
		winNumbers[utils.ParseInt(strings.TrimSpace(number))] = struct{}{}
	}

	numbers := make([]int, 0)
	for _, number := range cardLine[1] {
		if number == "" {
			continue
		}
		numbers = append(numbers, utils.ParseInt(number))
	}

	card := Card{
		WinNum:  winNumbers,
		Numbers: numbers,
	}
	return card
}

// part two
func part2(input string) int {
	cards := utils.SplitTo(input, "\n", parseCards)

	sum := 0
	freq := make([][]int, len(cards))

	for i, _ := range cards {
		CalcWins(i, &cards, &freq)
	}

	//freq2 := make([]int, len(cards), 0)
	for i := 0; i < len(freq); i++ {
		for j := 0; j < len(freq[i]); j++ {
			sum += freq[i][j]
		}
	}

	return sum
}

func CalcWins(cardNumber int, cards *[]Card, freq *[][]int) {
	if cardNumber >= len(*cards) {
		return
	}

	card := (*cards)[cardNumber]
	wins := card.GetNumberOfWinningNumbers()

	cardFreq := make([]int, len(*cards))
	if wins == 0 {
		cardFreq[cardNumber]++
		(*freq)[cardNumber] = cardFreq
		return
	}

	for i := cardNumber + 1; i <= cardNumber+wins && i < len(*cards); i++ {
		if (*freq)[i] == nil {
			CalcWins(i, cards, freq)
		}

		for k, count := range (*freq)[i] {
			cardFreq[k] += count
		}
	}

	cardFreq[cardNumber]++
	(*freq)[cardNumber] = cardFreq
}
