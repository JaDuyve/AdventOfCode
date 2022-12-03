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
	input, err := utils.ReadHTTP(2022, 03, session)
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
	rucksacks := strings.Split(input, "\n")
	var score int
	for _, rs := range rucksacks {
		item := findCommonItem(rs[:len(rs)/2], rs[len(rs)/2:])
		score += calcPriority(item)
	}
	return score
}

// part two
func part2(input string) int {
	rucksacks := strings.Split(input, "\n")
	var score int
	for i := 0; i < len(rucksacks); i += 3 {
		item := findCommonItem(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		score += calcPriority(item)
	}
	return score
}

func findCommonItem(items ...string) int32 {
	found := false

	for _, ch := range items[0] {
		for _, str := range items[1:] {
			found = strings.ContainsRune(str, ch)

			if found == false {
				break
			}
		}
		if found {
			return ch
		}
	}
	return -1
}

func calcPriority(ch rune) int {
	if ch >= 'a' && ch <= 'z' {
		return int(ch - ('a' - 1))
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch-('A'-1)) + 26
	}

	log.Fatalf("%c is a in valid priority character", ch)
	return -1
}
