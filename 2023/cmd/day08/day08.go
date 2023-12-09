package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"aoc/internal/utils"
)

var re = regexp.MustCompile(`(?m)(?P<key>[1-9A-Z]{3}) = \((?P<L>[1-9A-Z]{3}), (?P<R>[1-9A-Z]{3})\)`)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 8, session)
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
	splittedInput := strings.Split(input, "\n\n")
	directions := splittedInput[0]

	const endNodeKey = "ZZZ"
	currentNodeKey := "AAA"
	nodes := make(map[string]utils.Tuple[string, string])

	for _, match := range re.FindAllStringSubmatch(splittedInput[1], -1) {
		key := match[1]

		nodes[key] = utils.Tuple[string, string]{
			A: match[2],
			B: match[3],
		}
	}

	count := 0
	for currentNodeKey != endNodeKey {
		dir := directions[count%len(directions)]
		node := nodes[currentNodeKey]
		if dir == 'L' {
			currentNodeKey = node.A
		} else if dir == 'R' {
			currentNodeKey = node.B
		}

		count++
	}

	return count
}

// part two
func part2(input string) int {
	splittedInput := strings.Split(input, "\n\n")
	directions := splittedInput[0]

	var ghosts []Ghost
	nodes := make(map[string]utils.Tuple[string, string])

	for _, match := range re.FindAllStringSubmatch(splittedInput[1], -1) {
		key := match[1]

		nodes[key] = utils.Tuple[string, string]{
			A: match[2],
			B: match[3],
		}

		if key[2] == 'A' {
			ghosts = append(ghosts, Ghost{Key: key, Cycles: -1})
		}
	}

	steps := 0
	total := 1

	for len(ghosts) > 0 {
		dir := directions[steps%len(directions)]

		for i := len(ghosts) - 1; i >= 0; {
			current := ghosts[i].Key
			if current[len(current)-1] == 'Z' {
				total = lcm(total, steps)
				ghosts = utils.Remove(ghosts, i)
			} else {
				if dir == 'L' {
					ghosts[i].Key = nodes[current].A
				} else {
					ghosts[i].Key = nodes[current].B
				}
			}

			i--
		}

		steps++
	}

	return total
}

func gcd(a int, b int) int {
	for b != 0 {
		tmp := a
		a = b
		b = tmp % b
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

type Ghost struct {
	Cycles int
	Key    string
}

func AnyGhostNotAtEnd(ghosts *[]Ghost) bool {
	for _, g := range *ghosts {
		if g.Cycles == -1 {
			return true
		}
	}

	return false
}
