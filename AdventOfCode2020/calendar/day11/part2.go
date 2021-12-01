package day11

import (
	"AdventOfCode2020/utils/files"
)

func Part2() int {
	input := files.ReadFile("calendar/day11/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	grid := *new(SeatingGrid)
	grid.init(puzzleInput)

	first := true
	oldNumberOfOccupiedSeats, occupiedSeats := 0, 0

	for oldNumberOfOccupiedSeats != occupiedSeats || first {
		first = false
		grid.data = applyRulesPart2(grid)
		oldNumberOfOccupiedSeats = occupiedSeats
		occupiedSeats = grid.countOccupiedSeats()
	}

	return occupiedSeats
}

func applyRulesPart2(grid SeatingGrid) [][]rune {

	result := make([][]rune, len(grid.data))
	for i := 0; i < grid.rows; i++ {
		result[i] = append(result[i], grid.data[i]...)
	}

	for i := 1; i < grid.rows-1; i++ {
		for j := 1; j < grid.columns-1; j++ {

			switch grid.data[i][j] {
			case empty:
				if grid.countAdjacentOccupiedPart2(i, j) == 0 {
					result[i][j] = occupied
				}
			case occupied:
				if grid.countAdjacentOccupiedPart2(i, j) >= 5 {
					result[i][j] = empty
				}
			}
		}
	}

	return result
}

func (grid *SeatingGrid) countAdjacentOccupiedPart2(deltaY int, deltaX int) int {
	counter := 0

	for _, direction := range directions {
		first := true
		y, x := deltaY, deltaX

		for grid.data[y][x] == floor || first {
			first = false
			y += direction.deltaRow
			x += direction.deltaCol
			if grid.data[y][x] == occupied {
				counter++
				break
			}
		}
	}

	return counter
}
