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
	grid := strings.Split(input, "\n")
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isGearSymbol(grid[i][j]) == false {
				continue
			}

			numbers := getAllAttachedNumbers(j, i, &grid)

			if len(numbers) < 2 {
				continue
			}

			for n := 0; n < len(numbers); n++ {
				for m := n + 1; m < len(numbers); m++ {
					sum += numbers[n] * numbers[m]
				}
			}
		}
	}

	return sum
}

func isDigit(ch uint8) bool {
	return ch >= '0' && ch <= '9'
}

func isSymbol(ch uint8) bool {
	return !isDigit(ch) && ch != '.'
}

func isGearSymbol(ch uint8) bool {
	return ch == '*'
}

func getAllAttachedNumbers(x, y int, grid *[]string) []int {
	sby := utils.Max(y-1, 0)
	sbx := utils.Max(x-1, 0)
	eby := utils.Min(y+1, len(*grid)-1)
	ebx := utils.Min(x+1, len((*grid)[0])-1)
	numbers := make([]int, 0)
	coordinates := make(map[int]struct{})

	for i := sby; i <= eby; i++ {
		for j := sbx; j <= ebx; j++ {
			if isDigit((*grid)[i][j]) {
				sx, ex := getEntireNumberForCoordinates(j, i, grid)

				number := utils.ParseInt((*grid)[i][sx : ex+1])
				if _, ok := coordinates[number]; ok {
					continue
				}

				coordinates[number] = struct{}{}
				numbers = append(numbers, number)
			}
		}
	}

	return numbers
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

func getEntireNumberForCoordinates(x, y int, grid *[]string) (sx, ex int) {
	startX := x
	endX := x

	for i := x; i >= 0; i-- {
		if isDigit((*grid)[y][i]) {
			startX = i
		} else {
			i = -1
		}
	}

	for i := x; i < len((*grid)[y]); i++ {
		if isDigit((*grid)[y][i]) {
			endX = i
		} else {
			i = len((*grid)[y])
		}
	}

	return startX, endX
}
