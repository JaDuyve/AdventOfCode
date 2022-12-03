package main

import (
	"aoc/internal/utils"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	input, err := utils.ReadHTTP(2022, 01, os.Getenv("AOC_SESSION"))
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
	elfCalories := utils.SplitToIntGroups(input, "\n\n", "\n")
	sums := make([]int, len(elfCalories))
	for i, a := range elfCalories {
		sums[i] = utils.Sum[int](a)
	}

	return utils.Max(sums)
}

// part two
func part2(input string) int {
	elfCalories := utils.SplitToIntGroups(input, "\n\n", "\n")
	sums := make([]int, len(elfCalories))
	for i, a := range elfCalories {
		sums[i] = utils.Sum[int](a)
	}

	sort.Ints(sums)
	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}
