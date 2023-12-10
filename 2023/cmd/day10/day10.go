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
	input, err := utils.ReadHTTP(2023, 10, session)
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
	maze := strings.Split(input, "\n")

	x, y := FindStartCo(&maze)

	count := calcLoopCount(x, y, &maze)

	return count / 2
}

func calcLoopCount(sx, sy int, maze *[]string) int {
	count := 0

	px := -1
	py := -1
	x := sx
	y := sy
	for (*maze)[y][x] != 'S' || count == 0 {
		//if x < 0 || y < 0 || y >= len(*maze) || x >= len((*maze)[y]) {
		//	return 0
		//}
		//
		//if (*maze)[y][x] == '.' {
		//	return 0
		//}
		//
		//if isPossible(px, py, x, y, maze) == false {
		//	return 0
		//}

		count++
		pipe := (*maze)[y][x]
		tmpx := x
		tmpy := y
		x, y = nextStep(pipe, x, y, px, py)
		px = tmpx
		py = tmpy

	}

	return count
}

func isPossible(x, y, nx, ny int, maze *[]string) bool {
	if x == -1 || y == -1 {
		return true
	}

	currentPipe := (*maze)[y][x]
	nextPipe := (*maze)[ny][nx]

	switch currentPipe {
	case '|':
		if nextPipe == '-' {
			return false
		}

		if (nextPipe == '|' || nextPipe == 'F' || nextPipe == '7') && y > ny {
			return true
		}

		if (nextPipe == '|' || nextPipe == 'J' || nextPipe == 'L') && y < ny {
			return true
		}

		return false
		break
	case '-':
		if nextPipe == '|' {
			return false
		}

		if (nextPipe == '-' || nextPipe == '7' || nextPipe == 'J') && x < nx {
			return true
		}

		if (nextPipe == '-' || nextPipe == 'F' || nextPipe == 'L') && x > nx {
			return true
		}

		return false
		break
	case 'L':
		if (nextPipe == '|' || nextPipe == 'F' || nextPipe == '7') && y > ny {
			return true
		}

		if (nextPipe == '-' || nextPipe == '7' || nextPipe == 'J') && x < nx {
			return true
		}

		return false
		break
	case 'J':
		if (nextPipe == '|' || nextPipe == 'F' || nextPipe == '7') && y > ny {
			return true
		}

		if (nextPipe == '-' || nextPipe == 'F' || nextPipe == 'L') && x > nx {
			return true
		}

		return false
		break
	case '7':
		if (nextPipe == '|' || nextPipe == 'J' || nextPipe == 'L') && y < ny {
			return true
		}

		if (nextPipe == '-' || nextPipe == 'F' || nextPipe == 'L') && x > nx {
			return true
		}

		return false
		break
	case 'F':
		if (nextPipe == '|' || nextPipe == 'J' || nextPipe == 'L') && y < ny {
			return true
		}

		if (nextPipe == '-' || nextPipe == '7' || nextPipe == 'J') && x < nx {
			return true
		}

		return false
		break
	}

	return true
}

func nextStep(pipe uint8, x, y, px, py int) (nx, ny int) {
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
		x++
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
func part2(input string) int {
	return 0
}
