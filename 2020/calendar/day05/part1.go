package day05

import (
	"AdventOfCode2020/utils/files"
	"math"
	"strings"
)

type Seat struct {
	row    int
	column int
}

func (s *Seat) getSeatId() int {
	return s.row*8 + s.column
}

func Part1() int {
	input := files.ReadFile("calendar/day05/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	boardingPasses := strings.Split(puzzleInput, "\n")
	max := 0

	for _, boardingPass := range boardingPasses {
		seatId := newSeat(boardingPass).getSeatId()
		if seatId > max {
			max = seatId
		}
	}

	return max
}

func newSeat(boardingPass string) *Seat {
	seat := new(Seat)
	seat.row = calculatePosition(boardingPass[:7], 'F', 'B', 0, 127)
	seat.column = calculatePosition(boardingPass[7:], 'L', 'R', 0, 7)

	return seat
}

func calculatePosition(code string, lowerChar uint8, upperChar uint8, start int, end int) int {

	if len(code) == 1 {
		if code[0] == lowerChar {
			return start
		} else {
			return end
		}
	}

	if code[0] == lowerChar {
		end = start + (end-start)/2
		return calculatePosition(code[1:], lowerChar, upperChar, start, end)
	}

	if code[0] == upperChar {
		start += int(math.Ceil(float64(end-start) / 2))
		return calculatePosition(code[1:], lowerChar, upperChar, start, end)
	}

	return -1
}
