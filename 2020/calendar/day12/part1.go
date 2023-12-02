package day12

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part1() int {
	input := files.ReadFile("calendar/day12/input")
	return solvePart1(input)

}

type Coords struct {
	x int
	y int
}

func (c *Coords) move(value int, action string) {
	switch action {
	case "N":
		c.y += value
	case "E":
		c.x += value
	case "S":
		c.y -= value
	case "W":
		c.x -= value
	}
}

func solvePart1(puzzleInput string) int {
	ship := Coords{
		x: 0,
		y: 0,
	}
	maneuvers := strings.Split(puzzleInput, "\n")
	applyRoute(maneuvers, &ship)

	return aocmath.Abs(0-ship.x) + aocmath.Abs(0-ship.y)
}

func applyRoute(maneuvers []string, ship *Coords) {
	direction := 1

	for _, maneuver := range maneuvers {
		executeInstruction(maneuver[:1], conv.ToInt(maneuver[1:]), ship, &direction)
	}
}

func executeInstruction(action string, value int, ship *Coords, direction *int) {
	directions := []string{"N", "E", "S", "W"}

	switch action {
	case "L":
		*direction = (*direction - (value % 89) + 4) % 4
	case "R":
		*direction = (*direction + (value % 89)) % 4
	case "F":
		ship.move(value, directions[*direction])
	default:
		ship.move(value, action)
	}
}
