package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2022, 04, session)
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
	pairs := utils.SplitTo(input, "\n", ParseLine)
	count := 0

	for _, p := range pairs {
		if (p.A[0] >= p.B[0] && p.A[1] <= p.B[1]) || (p.A[0] <= p.B[0] && p.A[1] >= p.B[1]) {
			count++
		}
	}

	return count
}

// part two
func part2(input string) int {
	pairs := utils.SplitTo(input, "\n", ParseLine)
	count := 0

	for _, p := range pairs {
		if (p.A[0] >= p.B[0] && p.A[0] <= p.B[1]) || (p.A[1] >= p.B[0] && p.A[1] <= p.B[1]) ||
			(p.B[0] >= p.A[0] && p.B[0] <= p.A[1]) || (p.B[1] >= p.A[0] && p.B[1] <= p.A[1]) {
			count++
		}
	}

	return count
}

func ParseLine(data string) utils.Tuple[[]int, []int] {
	t := utils.Tuple[[]int, []int]{
		A: make([]int, 2),
		B: make([]int, 2),
	}

	_, err := fmt.Sscanf(data, "%d-%d,%d-%d", &t.A[0], &t.A[1], &t.B[0], &t.B[1])
	if err != nil {
		log.Fatal(err)
	}

	return t
}
