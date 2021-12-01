package day03

import (
	"log"
	"math"
)

func Run() {
	Part1()
}

func Part1() {

	solution := solveProblem1(325489)

	log.Printf("Day 03 part 3 solution: %d", solution)
}

func solveProblem1(input int) int {
	return calculateDistance(input)
}

func calculateCubeSize(number int) int {

	sqrtNumber := math.Sqrt(float64(number))

	size := int(math.Ceil(sqrtNumber))

	if size%2 == 0 {
		size += 1
	}

	return size
}

func calculateDistance(number int) int {
	cubeSize := calculateCubeSize(number)
	previousCubeSize := cubeSize - 2
	radius := cubeSize / 2

	maxPreviousCubeSize := previousCubeSize * previousCubeSize

	index := number - maxPreviousCubeSize

	height := (index % (cubeSize - 1)) - radius

	if height < 0 {
		height = -height
	}

	width := radius

	return width + height
}
