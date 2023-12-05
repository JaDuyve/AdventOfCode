package main

import (
	"cmp"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 05, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input)
	ms := time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res1, ms)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input)
	ms = time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res2, ms)
}

// part one
func part1(input string) int {
	mappingGroupStrings := utils.SplitToStringGroups(input, "\n\n", "\n")
	seeds := utils.SplitTo[int](strings.Split(mappingGroupStrings[0][0], ": ")[1], " ", utils.ParseInt)
	mappingGroups := make([][][]int, len(mappingGroupStrings)-1)

	for i, mg := range mappingGroupStrings[1:] {
		mappings := make([][]int, len(mg)-1)
		for j, m := range mg[1:] {
			mappings[j] = utils.SplitTo[int](m, " ", utils.ParseInt)
		}

		//slices.SortFunc(mappings, comp)
		mappingGroups[i] = mappings
	}

	min := math.MaxInt

	for _, seed := range seeds {
		currentValue := seed
		for i := 0; i < len(mappingGroups); i++ {
			for j := 0; j < len(mappingGroups[i]); j++ {
				m := mappingGroups[i][j]
				if m[1] <= currentValue && currentValue < m[1]+m[2] {
					currentValue = currentValue - m[1] + m[0]
					break
				}
			}
		}

		if currentValue < min {
			min = currentValue
		}
	}

	return min
}

func comp(a, b []int) int {
	return cmp.Compare(a[1], a[1])
}

// part two
func part2(input string) int {
	return 0
}
