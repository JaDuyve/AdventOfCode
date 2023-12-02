package grid

import (
	"AdventOfCode2021/util/aocstring"
	"strconv"
	"strings"
)

type Grid [][]int

func New(gridStr string) *Grid {
	rows := strings.Split(gridStr, "\n")
	grid := make(Grid, len(rows))

	for i, row := range rows {
		grid[i] = aocstring.SplitToIntArray(row, "")
	}

	return &grid
}

func (g *Grid) ApplyStep() int {
	g.IncreaseEnergyLevel()

	return g.Flash()
}

func (g *Grid) IncreaseEnergyLevel() {

	for y := range *g {
		for x := range (*g)[y] {
			(*g)[y][x]++
		}
	}
}

func (g *Grid) Flash() int {

	for !g.HaveAllOctopusesFlashed() {
		for y := range *g {
			for x := range (*g)[y] {
				if (*g)[y][x] == 10 {
					g.ApplyFlashOnCoordinate(y, x)
				}
			}
		}
	}

	counter := g.CountFlashedOctopuses()

	g.resetFlashedOctopus()

	return counter
}

func (g *Grid) ApplyFlashOnCoordinate(y, x int) {
	directions := [][]int{
		{0, -1},
		{-1, 0},
		{0, 1},
		{1, 0},
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}

	for _, direction := range directions {
		dirY := y + direction[0]
		dirX := x + direction[1]

		if dirY < 0 || dirX < 0 || dirY >= len(*g) || dirX >= len((*g)[y]) {
			continue
		}

		if (*g)[dirY][dirX] == 10 {
			continue
		}

		(*g)[dirY][dirX]++
	}

	(*g)[y][x]++
}

func (g Grid) HaveAllOctopusesFlashed() bool {
	for y := range g {
		for x := range g[y] {
			if g[y][x] == 10 {
				return false
			}
		}
	}

	return true
}

func (g Grid) ToString() string {
	str := ""

	for i, row := range g {
		if i != 0 {
			str += "\n"
		}

		for _, value := range row {
			str += strconv.Itoa(value)
		}

	}

	return str
}

func (g *Grid) resetFlashedOctopus() {
	for y := range *g {
		for x := range (*g)[y] {
			if (*g)[y][x] > 9 {
				(*g)[y][x] = 0
			}
		}
	}
}

func (g Grid) CountFlashedOctopuses() int {
	counter := 0

	for y := range g {
		for x := range g[y] {
			if g[y][x] > 9 {
				counter++
			}
		}
	}

	return counter
}

func (g Grid) IsSynchronised() bool {
	counter := 0

	for y := range g {
		for x := range g[y] {
			if g[y][x] == 0 {
				counter++
			}
		}
	}

	return counter == len(g)*len(g[0])
}
