package day04

import (
	"AdventOfCode2021/util/aocstring"
	"log"
	"regexp"
	"strings"
)

const (
	inputFile = "calendar/day04/input"
)

var expr = regexp.MustCompile("\\d+")

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 04 part 1: First Bingo !!!: %d", solvePart1(problem))
}

func solvePart1(bingoStr string) int {
	markings, boards := parseInput(bingoStr)

	for _, mark := range markings {
		applyMark(mark, &boards)

		boardIndex := checkIfWon(&boards)

		if boardIndex != -1 {
			return calculateSumUnmarked(boards[boardIndex]) * mark
		}
	}

	return 0
}

func applyMark(mark int, boards *[][][]int) {
	for i, board := range *boards {
		for j, row := range board {
			for k, val := range row {
				if val == mark {
					(*boards)[i][j][k] = -1
				}
			}
		}
	}
}

func checkIfWon(boards *[][][]int) int {
	for i := range *boards {
		if checkIfBoardWon(i, boards) {
			return i
		}
	}

	return -1
}

func checkIfBoardWon(boardIndex int, boards *[][][]int) bool {
	for _, row := range (*boards)[boardIndex] {
		counter := 0
		for _, val := range row {
			if val == -1 {
				counter++
			} else {
				break
			}
		}

		if counter == 5 {
			return true
		}

	}

	for j := range (*boards)[boardIndex] {
		counter := 0
		for k := range (*boards)[boardIndex][j] {
			if (*boards)[boardIndex][k][j] == -1 {
				counter++
			} else {
				break
			}
		}

		if counter == 5 {
			return true
		}

	}

	return false
}

func calculateSumUnmarked(board [][]int) int {
	sum := 0

	for _, row := range board {
		for _, val := range row {
			if val != -1 {
				sum += val
			}
		}
	}

	return sum
}

func parseInput(input string) ([]int, [][][]int) {
	splitInput := strings.Split(input, "\n\n")

	markings := aocstring.SplitToIntArray(splitInput[0], ",")
	boards := parseBoards(splitInput[1:])

	return markings, boards
}

func parseBoards(boardStrings []string) [][][]int {
	boards := make([][][]int, len(boardStrings))

	for i, boardString := range boardStrings {
		boards[i] = parseBoard(boardString)
	}

	return boards
}

func parseBoard(boardStr string) [][]int {
	rows := strings.Split(boardStr, "\n")
	board := make([][]int, len(rows))

	for i, row := range rows {
		a := expr.FindAllString(row, -1)
		board[i] = aocstring.StringArrayToIntArray(a)
	}

	return board
}

func part2(problem string) {
	log.Printf("Day 04 part 2: Last Bingo !!!: %d", solvePart2(problem))
}

func solvePart2(bingoStr string) int {
	markings, boards := parseInput(bingoStr)
	lastScore := 0

	for _, mark := range markings {
		applyMark(mark, &boards)

		boardIndex := checkIfWon(&boards)

		if boardIndex != -1 {
			lastScore = calculateSumUnmarked(boards[boardIndex]) * mark

			for i := len(boards) - 1; i >= 0; i-- {
				if checkIfBoardWon(i, &boards) {
					copy(boards[i:], boards[i+1:])
					boards[len(boards)-1] = nil
					boards = boards[:len(boards)-1]
				}
			}

			if len(boards) == 0 {
				return lastScore
			}
		}
	}

	return lastScore
}
