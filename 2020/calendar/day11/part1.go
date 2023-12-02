package day11

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"strings"
)

type SeatingGrid struct {
	columns int
	rows    int
	data    [][]rune
}

type Direction struct {
	deltaRow int
	deltaCol int
}

var directions = []Direction{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
}

const (
	occupied = '#'
	empty    = 'L'
	floor    = '.'
)

func (grid *SeatingGrid) countOccupiedSeats() int {
	occupiedSeats := 0

	for row := 0; row < grid.rows; row++ {
		for column := 0; column < grid.columns; column++ {
			if grid.data[row][column] == '#' {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func (grid *SeatingGrid) print() {
	slices.Print2DimRuneArray(grid.data)
	println()
}

func (grid *SeatingGrid) init(input string) {
	inputArray := strings.Split(input, "\n")
	frameCharacter := "="

	frame := make([][]rune, 0, len(inputArray)+2)
	padding := strings.Repeat(frameCharacter, len(inputArray[0])+2)

	frame = append(frame, []rune(padding))

	for _, row := range inputArray {
		frame = append(frame, []rune(frameCharacter+row+frameCharacter))
	}

	frame = append(frame, []rune(padding))

	grid.data = frame
	grid.rows = len(frame)
	grid.columns = len(frame[0])
}

func Part1() int {
	input := files.ReadFile("calendar/day11/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	grid := *new(SeatingGrid)
	grid.init(puzzleInput)

	first := true
	oldNumberOfOccupiedSeats, occupiedSeats := 0, 0
	for oldNumberOfOccupiedSeats != occupiedSeats || first {
		first = false
		grid.data = applyRules(grid)
		oldNumberOfOccupiedSeats = occupiedSeats
		occupiedSeats = grid.countOccupiedSeats()
	}

	return occupiedSeats
}

func applyRules(grid SeatingGrid) [][]rune {

	result := make([][]rune, len(grid.data))
	for i := 0; i < grid.rows; i++ {
		result[i] = append(result[i], grid.data[i]...)
	}

	for i := 1; i < grid.rows-1; i++ {
		for j := 1; j < grid.columns-1; j++ {

			switch grid.data[i][j] {
			case empty:
				if grid.countAdjacentOccupied(i, j) == 0 {
					result[i][j] = occupied
				}
			case occupied:
				if grid.countAdjacentOccupied(i, j) >= 4 {
					result[i][j] = empty
				}
			}
		}
	}

	return result
}

func (grid *SeatingGrid) countAdjacentOccupied(deltaY int, deltaX int) int {
	counter := 0

	for _, direction := range directions {
		if grid.data[deltaY+direction.deltaRow][deltaX+direction.deltaCol] == occupied {
			counter++
		}
	}

	return counter
}
