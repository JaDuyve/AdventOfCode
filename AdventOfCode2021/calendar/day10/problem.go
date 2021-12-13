package day10

import (
	"AdventOfCode2021/util/aocstring"
	"log"
	"strings"
)

const (
	inputFile = "calendar/day09/input"
)

var opposites = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
}

func part1(problem string) {
	log.Printf("Day 09 part 1: Smoke Basin: %d", solvePart1(problem))
}

func solvePart1(problem string) int {
	subsystem := strings.Split(problem, "\n")

	pointArray := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	freq := make(map[byte]int)

	for _, chunk := range subsystem {
		if ok, bracket := isChunkValid(chunk); !ok {
			freq[bracket]++
		}
	}

	points := 0
	for key, c := range freq {
		points += c * pointArray[key]
	}

	return points
}

func filterSyntaxErrors(subsystem []string) {
	for i := len(subsystem) - 1; i >= 0; i-- {
		if ok, _ := isChunkValid(subsystem[i]); !ok {
			subsystem[i] = subsystem[len(subsystem)-1]
			subsystem[len(subsystem)-1] = ""
			subsystem = subsystem[:len(subsystem)-1]
		}
	}
}

func isChunkValid(chunk string) (bool, byte) {
	freq := map[byte]int{
		'(': 0,
		'[': 0,
		'{': 0,
		'<': 0,
	}

	for _, c := range chunk {
		bc := byte(c)
		if bc == '(' {
			freq['(']++
		} else if bc == ')' && freq['('] == 0 {
			return false, ')'
		} else if bc == ')' {
			freq['(']--
		}

		if bc == '[' {
			freq['[']++
		} else if bc == ']' && freq['['] == 0 {
			return false, ']'
		} else if bc == ']' {
			freq['[']--
		}

		if bc == '{' {
			freq['{']++
		} else if bc == '}' && freq['{'] == 0 {
			return false, '}'
		} else if bc == '}' {
			freq['{']--
		}

		if bc == '<' {
			freq['<']++
		} else if bc == '>' && freq['<'] == 0 {
			return false, '>'
		} else if bc == '>' {
			freq['<']--
		}
	}

	for key, value := range freq {
		if value != 0 {
			return false, opposites[key]
		}
	}

	return true, '.'
}
