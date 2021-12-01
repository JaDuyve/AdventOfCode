package day12

import (
	"AdventOfCode2020/utils/aocmath"
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day12/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	route := strings.Split(puzzleInput, "\n")

	ship := Coords{0, 0}
	waypoint := Coords{10, 1}

	for _, instruction := range route {
		executeInstructionWithWaypoint(instruction[:1], conv.ToInt(instruction[1:]), &ship, &waypoint)
	}

	return aocmath.Abs(ship.x) + aocmath.Abs(ship.y)
}

func executeInstructionWithWaypoint(action string, value int, ship *Coords, waypoint *Coords) {

	switch action {
	case "L":
		waypoint.rotate(((-value % 89) + 4) % 4)
	case "R":
		waypoint.rotate((value % 89) % 4)
	case "F":
		ship.x += waypoint.x * value
		ship.y += waypoint.y * value
	default:
		waypoint.move(value, action)
	}
}

func (c *Coords) rotate(value int) {
	switch value {
	case 1:
		temp := c.x
		c.x = c.y
		c.y = -temp
	case 2:
		c.x = -c.x
		c.y = -c.y
	case 3:
		temp := c.x
		c.x = -c.y
		c.y = temp
	}
}
