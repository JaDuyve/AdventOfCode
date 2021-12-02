package day02

import (
	"AdventOfCode2021/calendar/day02/direction"
	"io/ioutil"
	"log"
	"strings"
)

const (
	inputFilePart1 = "calendar/day02/input"
)

func Run() {
	part1()
	//part2()
}

func part1() {
	problem := readFile(inputFilePart1)

	log.Printf("Day 02 part 1: Calculated distance: %d", solvePart1(problem))
}

func solvePart1(pathString string) int {
	path := strings.Split(pathString, "\n")

	x, y := calculateEndCoordinate(path)

	return x * y
}

func calculateEndCoordinate(commandStrings []string) (x, y int) {

	for _, commandString := range commandStrings {
		command, err := direction.NewCommand(commandString)
		if err != nil {
			log.Fatalf("commandString is not valid, commandString: '%s', error: '%v'", commandString, err)
		}

		switch command.Direction {
		case direction.FORWARD:
			x += command.Distance
			break
		case direction.DOWN:
			y += command.Distance
			break
		case direction.UP:
			y -= command.Distance
			break
		}
	}

	return
}

func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
