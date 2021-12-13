package day10

import (
	"AdventOfCode2021/util/aocstring"
	"AdventOfCode2021/util/stack"
	"log"
	"sort"
	"strings"
)

const (
	inputFile = "calendar/day10/input"
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
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 10 part 1: Syntax Scoring: %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 10 part 2: Syntax Scoring: %d", solvePart2(problem))
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

func filterSyntaxErrors(subsystem *[]string) {
	for i := len(*subsystem) - 1; i >= 0; i-- {
		if ok, _ := isChunkValid((*subsystem)[i]); !ok {
			copy((*subsystem)[i:], (*subsystem)[i+1:])
			(*subsystem)[len(*subsystem)-1] = ""
			*subsystem = (*subsystem)[:len(*subsystem)-1]
		}
	}
}

func isChunkValid(chunk string) (bool, byte) {
	st := new(stack.Stack)

	beginBrackets := make(map[byte]struct{})
	beginBrackets['('] = struct{}{}
	beginBrackets['{'] = struct{}{}
	beginBrackets['['] = struct{}{}
	beginBrackets['<'] = struct{}{}

	for _, c := range chunk {
		bc := byte(c)

		if _, ok := beginBrackets[bc]; ok {
			st.Push(bc)
		} else {
			d, k := st.Pop()

			if !k {
				return false, bc
			}

			if bc != opposites[d] {
				return false, bc
			}
		}
	}

	return true, '.'
}

func solvePart2(problem string) int {
	subsystem := strings.Split(problem, "\n")

	filterSyntaxErrors(&subsystem)

	scores := make([]int, 0)

	for _, chunk := range subsystem {
		score := calculatePointsForChunk(chunk)

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[(len(scores) / 2)]
}

func calculatePointsForChunk(chunk string) int {
	pointArray := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	completionArr := fixChunk(chunk)

	score := 0

	for _, ch := range completionArr {
		score *= 5

		score += pointArray[ch]
	}

	return score
}

func fixChunk(chunk string) []byte {
	comp := make([]byte, 0)
	st := new(stack.Stack)

	beginBrackets := make(map[byte]struct{})
	beginBrackets['('] = struct{}{}
	beginBrackets['{'] = struct{}{}
	beginBrackets['['] = struct{}{}
	beginBrackets['<'] = struct{}{}

	for _, c := range chunk {
		bc := byte(c)

		if _, ok := beginBrackets[bc]; ok {
			st.Push(bc)
		} else {
			d, k := st.Pop()

			if !k {
				log.Fatalf("chunk is invalid at %c", bc)
			}

			if bc != opposites[d] {
				log.Fatalf("chunk is invalid at %c", bc)
			}
		}
	}

	for true {
		d, ok := st.Pop()
		if !ok {
			break
		}

		comp = append(comp, opposites[d])
	}

	return comp
}
