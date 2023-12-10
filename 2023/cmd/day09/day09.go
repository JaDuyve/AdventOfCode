package main

import (
	"log"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 9, session)
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
	histories := utils.SplitToIntGroups(input, "\n", " ")

	sum := 0
	for i := 0; i < len(histories); i++ {
		sum += calcHistoryValue(histories[i])
	}
	return sum
}

func calcHistoryValue(history []int) int {
	comp := func(i int) bool {
		return history[i] != 0
	}

	count := 0
	for utils.Any(history[:len(history)-count], comp) {
		for i := 0; i < len(history)-count-1; i++ {
			history[i] = history[i+1] - history[i]
		}

		count++
	}

	sum := 0
	for i := count - 1; i >= 0; i-- {
		sum += history[len(history)-i-1]
	}

	return sum
}

// part two
func part2(input string) int {
	return 0
}
