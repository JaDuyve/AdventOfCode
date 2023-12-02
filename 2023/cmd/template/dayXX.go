package main

import (
	"log"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 01, session)
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
	return 0
}

// part two
func part2(input string) int {
	return 0
}
