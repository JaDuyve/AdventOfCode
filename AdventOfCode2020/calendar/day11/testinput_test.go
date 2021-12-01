package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CountAdjacentOccupiedPart2_8AdjacentSeats(t *testing.T) {
	input := ".......#.\n...#.....\n.#.......\n.........\n..#L....#\n....#....\n.........\n#........\n...#....."
	grid := *new(SeatingGrid)
	grid.init(input)

	expected := 8
	assert.Equal(t, expected, grid.countAdjacentOccupiedPart2(5, 4))
}

func Test_CountAdjacentOccupiedPart2_NoAdjacentSeats(t *testing.T) {
	input := ".##.##.\n#.#.#.#\n##...##\n...L...\n##...##\n#.#.#.#\n.##.##."
	grid := *new(SeatingGrid)
	grid.init(input)

	expected := 0
	assert.Equal(t, expected, grid.countAdjacentOccupiedPart2(4, 4))
}

func Test_CountAdjacentOccupiedPart2_AdjacentEmptySeat(t *testing.T) {
	input := ".............\n.L.L.#.#.#.#.\n............."
	grid := *new(SeatingGrid)
	grid.init(input)

	expected := 0
	assert.Equal(t, expected, grid.countAdjacentOccupiedPart2(2, 2))
}

func Test_SolvePart1_TestInput(t *testing.T) {
	input := "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	expected := 37
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInput(t *testing.T) {
	input := "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	expected := 26
	assert.Equal(t, expected, solvePart2(input))
}
