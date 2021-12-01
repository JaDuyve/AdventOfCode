package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalculatePosition_Row(t *testing.T) {
	boardingPassRowsAnswerPairs := map[string]int{"FBFBBFF": 44, "BFFFBBF": 70, "FFFBBBF": 14, "BBFFBBF": 102}

	for code, answer := range boardingPassRowsAnswerPairs {
		assert.Equalf(t,
			answer,
			calculatePosition(code, 'F', 'B', 0, 127),
			"Answer for %s should be %d.",
			code,
			answer)
	}
}

func Test_CalculatePosition_Column(t *testing.T) {
	boardingPassRowsAnswerPairs := map[string]int{"RLR": 5, "RRR": 7, "RLL": 4}

	for code, answer := range boardingPassRowsAnswerPairs {
		assert.Equalf(t,
			answer,
			calculatePosition(code, 'L', 'R', 0, 7),
			"Answer for %s should be %d.",
			code,
			answer)
	}
}

func Test_CalculateSeatId(t *testing.T) {
	boardingPassSeatIdPairs := map[string]int{"FBFBBFFRLR": 357, "BFFFBBFRRR": 567, "FFFBBBFRRR": 119, "BBFFBBFRLL": 820}

	for boardingPass, seatId := range boardingPassSeatIdPairs {
		assert.Equalf(t,
			seatId,
			newSeat(boardingPass).getSeatId(),
			"Answer for %s should be %d.",
			boardingPass,
			seatId)
	}
}
