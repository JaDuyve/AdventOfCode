package main

import (
	"aoc/internal/utils"
	"log"
	"os"
)

func main() {
	input, err := utils.ReadHTTP(2022, 01, os.Getenv("AOC_SESSION"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	log.Println("Result:", part1(input))
	log.Println("--- Part Two ---")
	log.Println("Result:", part2(input))
}

// part one
func part1(input string) int {
	return 0
}

// part two
func part2(input string) int {
	return 0
}
