package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"aoc/internal/utils"
)

var (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 02, session)
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

type Game struct {
	r int
	g int
	b int
}

func (g Game) IsPossible(g2 Game) bool {
	return g.r <= g2.r && g.g <= g2.g && g.b <= g2.b
}

func (g Game) Value() int {
	return g.r * g.g * g.b
}

var handPattern = regexp.MustCompile(" (?P<count>\\d+) (?P<type>\\w+)")

// part one
func part1(input string) int {
	maxGame := Game{
		r: 12,
		g: 13,
		b: 14,
	}

	games := utils.SplitTo(input, "\n", parseGame)

	sum := 0

	for i, game := range games {
		if game.IsPossible(maxGame) {
			sum += i + 1
		}
	}

	return sum
}

func parseGame(line string) Game {
	game := Game{}

	sets := utils.SplitToStringGroups(strings.Split(line, ":")[1], ";", ",")

	for _, set := range sets {
		for _, hand := range set {
			x := handPattern.FindStringSubmatch(hand)
			count := utils.ParseInt(x[1])
			switch x[2] {
			case RED:
				game.r = utils.Max(game.r, count)
				break
			case GREEN:
				game.g = utils.Max(game.g, count)
				break
			case BLUE:
				game.b = utils.Max(game.b, count)
				break
			}
		}
	}

	return game
}

// part two
func part2(input string) int {
	games := utils.SplitTo(input, "\n", parseGame)

	sum := 0

	for _, game := range games {
		sum += game.Value()
	}

	return sum
}
