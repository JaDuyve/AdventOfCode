package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

type Move struct {
	from  int
	to    int
	count int
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2022, 05, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input)
	ms := time.Since(startTime).Milliseconds()
	log.Printf("Result: %s in %dms\n", res1, ms)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input)
	ms = time.Since(startTime).Milliseconds()
	log.Printf("Result: %s in %dms\n", res2, ms)
}

// part one
func part1(input string) string {
	state, moves := Parse(input)

	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			tmp, _ := state[m.from-1].Pop()
			state[m.to-1].Push(tmp)
		}
	}

	res := strings.Builder{}
	for _, s := range state {
		tmp, _ := s.Top()
		res.WriteRune(tmp)
	}
	return res.String()
}

// part two
func part2(input string) string {
	state, moves := Parse(input)

	for _, m := range moves {
		tmpStack := utils.Stack[rune]{}
		for i := 0; i < m.count; i++ {
			tmp, _ := state[m.from-1].Pop()
			tmpStack.Push(tmp)
		}
		for !tmpStack.IsEmpty() {
			tmp, _ := tmpStack.Pop()
			state[m.to-1].Push(tmp)
		}
	}

	res := strings.Builder{}
	for _, s := range state {
		tmp, _ := s.Top()
		res.WriteRune(tmp)
	}
	return res.String()
}

func Parse(input string) ([]utils.Stack[rune], []Move) {
	inputParts := strings.Split(input, "\n\n")
	stateString := inputParts[0]
	movesString := inputParts[1]

	state := ParseState(stateString)
	moves := utils.SplitTo(movesString, "\n", ParseMove)

	return state, moves
}

func ParseState(data string) []utils.Stack[rune] {
	currentStateString := TransposeString(data)

	currentStateString = strings.Replace(currentStateString, "[", "", -1)
	currentStateString = strings.Replace(currentStateString, "]", "", -1)
	currentStateString = strings.Replace(currentStateString, " ", "", -1)
	currentStateString = strings.Replace(currentStateString, "\n\n\n", "", -1)
	currentStateString = strings.TrimSpace(currentStateString)

	l := strings.Split(currentStateString, "\n")

	stacks := make([]utils.Stack[rune], len(l))
	for i, val := range l {
		st := utils.NewStack[rune]()
		for _, r := range val[1:] {
			st.Push(r)
		}

		stacks[i] = *st
	}

	return stacks
}

func ParseMove(data string) Move {
	move := Move{}

	_, err := fmt.Sscanf(data, "move %d from %d to %d", &move.count, &move.from, &move.to)
	if err != nil {
		log.Fatal(err)
	}

	return move
}

func TransposeString(data string) string {
	tmp := strings.Split(data, "\n")
	stacks := make([]string, len(tmp[0]))

	for j := 0; j < len(tmp[0]); j++ {
		str := strings.Builder{}
		for i := len(tmp) - 1; i >= 0; i-- {
			str.WriteByte(tmp[i][j])
		}

		stacks[j] = str.String()
	}

	return strings.Join(stacks, "\n")
}
