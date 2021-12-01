package day01

import (
	"io/ioutil"
	"log"
)

const inputFilePart1 = "day01/input1"

func Run() {
	part1()
	part2()
}

func part1() {
	problem := readFile(inputFilePart1)

	log.Printf("Day 01 part 1: Destination floor: %d", solveProblem1(problem))
}

func solveProblem1(problem string) int {
	calculatedFloor := calculateFloor(problem, 0, 0)

	return calculatedFloor
}

func calculateFloor(floorPath string, currentFloor int, currentPosition int) int {

	if len(floorPath) <= currentPosition {
		return currentFloor
	}

	bracket := floorPath[currentPosition]

	if bracket == '(' {
		currentFloor++
	} else if bracket == ')' {
		currentFloor--
	}

	return calculateFloor(floorPath, currentFloor, currentPosition+1)
}

func part2() {
	problem := readFile(inputFilePart1)

	log.Printf("Day 01 part 2: Position of basement: %d", solveProblem2(problem))
}

func solveProblem2(problem string) int {
	positionBasement := calculatePositionForFloor(problem, -1, 0, 0)

	return positionBasement
}

func calculatePositionForFloor(floorPath string, wantedFloor int, currentFloor int, currentPosition int) int {
	if len(floorPath) <= currentPosition {
		return currentPosition
	}

	if wantedFloor == currentFloor {
		return currentPosition
	}

	bracket := floorPath[currentPosition]

	if bracket == '(' {
		currentFloor++
	} else if bracket == ')' {
		currentFloor--
	}

	return calculatePositionForFloor(floorPath, wantedFloor, currentFloor, currentPosition+1)
}

func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
