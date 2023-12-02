package day05

import (
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day05/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	airplane := make([][]int, 128)
	boardingPasses := strings.Split(puzzleInput, "\n")

	for _, boardingPass := range boardingPasses {
		seat := newSeat(boardingPass)
		if airplane[seat.row] == nil {
			airplane[seat.row] = make([]int, 8)
		}
		airplane[seat.row][seat.column] = seat.getSeatId()
	}

	return searchSeat(airplane)
}

func searchSeat(airplane [][]int) int {

	for i := 0; i < len(airplane); i++ {
		if airplane[i] == nil {
			continue
		}

		for j := 1; j < len(airplane[i])-1; j++ {
			if airplane[i][j] == 0 && airplane[i][j+1]-airplane[i][j-1] == 2 {
				return airplane[i][j-1] + 1
			}
		}
	}

	return -1
}
