package main

import (
	"aoc/internal/utils"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 03, session)
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
	grid := strings.Split(input, "\n")
	//grid2 := grid
	sum := 0

	for i := 0; i < len(grid); i++ {
		startIndexNumber := -1
		for j := 0; j < len(grid[i]); j++ {
			if startIndexNumber == -1 && isDigit(grid[i][j]) {
				startIndexNumber = j
			}

			if startIndexNumber != -1 && (j+1 == len(grid[i]) || isDigit(grid[i][j+1]) == false) {
				if isNextToSymbol(i, startIndexNumber, i, j, &grid) {
					sum += utils.ParseInt(grid[i][startIndexNumber : j+1])
				}
				startIndexNumber = -1
			}
		}
	}

	return sum
}

// part two
func part2(input string) int {
	return 0
}

func isDigit(ch uint8) bool {
	return ch >= '0' && ch <= '9'
}

func isSymbol(ch uint8) bool {
	return !isDigit(ch) && ch != '.'
}

func isNextToSymbol(sy, sx, ey, ex int, grid *[]string) bool {
	sby := utils.Max(sy-1, 0)
	sbx := utils.Max(sx-1, 0)
	eby := utils.Min(ey+1, len(*grid)-1)
	ebx := utils.Min(ex+1, len((*grid)[0])-1)

	for i := sby; i <= eby; i++ {
		for j := sbx; j <= ebx; j++ {
			if isSymbol((*grid)[i][j]) {
				return true
			}
		}
	}
	return false
}
