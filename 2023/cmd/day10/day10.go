package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 10, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input, '-')
	us := time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res1, us)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input, '-')
	us = time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res2, us)
}

// part one
func part1(input string, startChar uint8) int {
	maze := strings.Split(input, "\n")

	x, y := FindStartCo(&maze)

	count := calcLoopCount(x, y, &maze, startChar)

	return count / 2
}

func calcLoopCount(sx, sy int, maze *[]string, startChar uint8) int {
	count := 0

	px := -1
	py := -1
	x := sx
	y := sy
	for (*maze)[y][x] != 'S' || count == 0 {
		count++
		pipe := (*maze)[y][x]
		tmpx := x
		tmpy := y
		x, y = nextStep(pipe, x, y, px, py, startChar)
		px = tmpx
		py = tmpy

	}

	return count
}

func nextStep(pipe uint8, x, y, px, py int, startChar uint8) (nx, ny int) {
	if pipe == 'S' {
		pipe = startChar
	}

	switch pipe {
	case '|':
		if y-1 != py {
			y--
		} else {
			y++
		}
		break
	case '-':
		if x-1 != px {
			x--
		} else {
			x++
		}
		break
	case 'L':
		if x+1 != px {
			x++
		} else {
			y--
		}
		break
	case 'J':
		if x-1 != px {
			x--
		} else {
			y--
		}
		break
	case '7':
		if x-1 != px {
			x--
		} else {
			y++
		}
		break
	case 'F':
		if x+1 != px {
			x++
		} else {
			y++
		}
		break
	case 'S':
		x--
		break
	}

	return x, y
}

func FindStartCo(maze *[]string) (int, int) {
	for y := 0; y < len(*maze); y++ {
		for x := 0; x < len((*maze)[y]); x++ {
			if (*maze)[y][x] == 'S' {
				return x, y
			}
		}
	}

	return -1, -1
}

// part two
func part2(input string, startChar uint8) int {
	maze := strings.Split(input, "\n")

	sx, sy := FindStartCo(&maze)

	m := calcLoopCount2(sx, sy, &maze, startChar)

	maze[sy] = strings.ReplaceAll(maze[sy], "S", string(startChar))
	count := 0

	for y := 0; y < len(maze); y++ {
		count += CalcInClosedTilesForRow(y, maze[y], &m)
	}

	return count
}

func CalcInClosedTilesForRow(y int, row string, m *map[string]struct{}) int {
	count := 0

	isUp := 0
	for x := 0; x < len(row); x++ {
		if _, ok := (*m)[fmt.Sprintf("%d-%d", x, y)]; ok {
			tile := row[x]
			if tile == '|' {
				isUp++
				continue
			}

			prevTile := tile

			for xx := x + 1; xx < len(row); xx++ {
				currTile := row[xx]

				if currTile != '-' {

					if prevTile == 'F' && currTile == 'J' || prevTile == 'L' && currTile == '7' {
						isUp++
					}

					x = xx
					break
				}
			}

		} else if isUp%2 != 0 {
			count++
		}

	}
	return count
}

func calcLoopCount2(sx, sy int, maze *[]string, startChar uint8) map[string]struct{} {
	count := 0
	m := make(map[string]struct{})

	px := -1
	py := -1
	x := sx
	y := sy
	m[fmt.Sprintf("%d-%d", x, y)] = struct{}{}
	for (*maze)[y][x] != 'S' || count == 0 {
		m[fmt.Sprintf("%d-%d", x, y)] = struct{}{}

		count++
		pipe := (*maze)[y][x]
		tmpx := x
		tmpy := y
		x, y = nextStep(pipe, x, y, px, py, startChar)
		px = tmpx
		py = tmpy

	}
	m[fmt.Sprintf("%d-%d", x, y)] = struct{}{}

	return m
}
