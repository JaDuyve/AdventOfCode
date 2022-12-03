package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

var attributePoints = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissor
}

var moves = map[string]string{
	"AX": "Z",
	"AY": "X",
	"AZ": "Y",
	"BX": "X",
	"BY": "Y",
	"BZ": "Z",
	"CX": "Y",
	"CY": "Z",
	"CZ": "X",
}

var scores = map[string]int{
	"AX": 3, // rock rock -> draw
	"AY": 6, // rock Paper -> won
	"AZ": 0, // rock Scissor -> lost
	"BX": 0, // Paper rock -> lost
	"BY": 3, // Paper Paper -> draw
	"BZ": 6, // Paper Scissor -> win
	"CX": 6, // Scissor rock -> won
	"CY": 0, // Scissor Paper -> lost
	"CZ": 3, // Scissor Scissor -> draw
}

type Round struct {
	playerA string
	playerB string
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2022, 02, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input)
	ms := time.Since(startTime).Milliseconds()
	log.Printf("Result: %d in %dms\n", res1, ms)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input)
	ms = time.Since(startTime).Milliseconds()
	log.Printf("Result: %d in %dms\n", res2, ms)
}

// part one
func part1(input string) int {
	rounds := parseToRounds(input)
	score := 0

	for _, round := range rounds {
		r := round.playerA + round.playerB
		score += attributePoints[round.playerB] + scores[r]
	}

	return score
}

// part two
func part2(input string) int {
	rounds := parseToRounds(input)
	score := 0

	for _, round := range rounds {
		r := round.playerA + round.playerB

		move := moves[r]
		score += attributePoints[move] + scores[round.playerA+move]
	}

	return score
}

func parseToRounds(input string) []Round {
	roundStrings := strings.Split(input, "\n")

	rounds := make([]Round, len(roundStrings))

	for i, val := range roundStrings {
		round := Round{}
		_, err := fmt.Sscanf(val, "%s %s", &round.playerA, &round.playerB)
		if err != nil {
			log.Fatal(err)
		}
		rounds[i] = round
	}
	return rounds
}
