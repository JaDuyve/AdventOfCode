package main

import (
	"log"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2022, 06, session)
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
	const markerSize = 4
	return CalcMarkerPosition(input, markerSize)
}

// part two
func part2(input string) int {
	const markerSize = 14
	return CalcMarkerPosition(input, markerSize)
}

func CalcMarkerPosition(input string, markerSize int) int {
	count := markerSize

	for i := markerSize; i < len(input); i++ {
		freq := make(map[uint8]bool)
		for j := i - markerSize; j < i; j++ {
			if ok, _ := freq[input[j]]; !ok {
				freq[input[j]] = true
			} else {
				break
			}
		}

		if len(freq) == markerSize {
			break
		} else {
			count++
		}
	}

	return count
}
