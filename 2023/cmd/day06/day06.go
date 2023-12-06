package main

import (
	"log"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 06, session)
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

// part one
func part1(input string) int {
	groups := utils.SplitToStringGroups(input, "\n", ":")

	times := utils.SplitToIgnoreSpace[int](groups[0][1], " ", utils.ParseInt)
	distances := utils.SplitToIgnoreSpace[int](groups[1][1], " ", utils.ParseInt)

	factor := 1

	for i := 0; i < len(times); i++ {
		count := 0
		time := times[i]
		recordDistance := distances[i]

		for j := 1; j < time; j++ {
			distance := j * (time - j)
			if distance > recordDistance {
				count++
			}
		}

		factor *= count
	}

	return factor
}

// part two
func part2(input string) int {
	groups := utils.SplitToStringGroups(input, "\n", ":")

	times := utils.SplitToIgnoreSpace[int](strings.ReplaceAll(groups[0][1], " ", ""), " ", utils.ParseInt)
	distances := utils.SplitToIgnoreSpace[int](strings.ReplaceAll(groups[1][1], " ", ""), " ", utils.ParseInt)

	factor := 1

	for i := 0; i < len(times); i++ {
		count := 0
		time := times[i]
		recordDistance := distances[i]

		for j := 1; j < time; j++ {
			distance := j * (time - j)
			if distance > recordDistance {
				count++
			}
		}

		factor *= count
	}

	return factor
}
