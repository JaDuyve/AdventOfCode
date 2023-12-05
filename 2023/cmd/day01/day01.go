package main

import (
	"aoc/internal/utils"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var numberRegex = regexp.MustCompile("[a-zA-Z]")

func main() {
	input, err := utils.ReadHTTP(2023, 01, os.Getenv("AOC_SESSION"))
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
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		sum += getNumberValueForLine(line)
	}

	return sum
}

func getNumberValueForLine(line string) int {
	cleanedLine := numberRegex.ReplaceAllString(line, "")
	return utils.ParseInt(cleanedLine[0:1] + cleanedLine[len(cleanedLine)-1:])
}

// part two
func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		sum += getNumberValueForLine2(line)
	}

	return sum
}

var numberStrings = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func getNumberValueForLine2(line string) int {
	for key, value := range numberStrings {
		line = strings.ReplaceAll(line, key, value)
	}

	cleanedLine := numberRegex.ReplaceAllString(line, "")
	return utils.ParseInt(cleanedLine[0:1] + cleanedLine[len(cleanedLine)-1:])
}
