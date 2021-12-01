package day01

import (
	"AdventOfCode2021/util/aocstring"
	"io/ioutil"
	"log"
)

const inputFilePart1 = "calendar/day01/input"

func Run() {
	part1()
	part2()
}

func part1() {
	problem := readFile(inputFilePart1)

	log.Printf("Day 01 part 1: Number of measurement increases: %d", calculateNumberOfIncreases(problem))
}

func calculateNumberOfIncreases(measurementString string) int {

	measurements := aocstring.SplitToIntArray(measurementString, "\n")

	counter := 0

	currentMeasurement := measurements[0]

	for _, measurement := range measurements[1:] {
		nextMeasurement := measurement

		if nextMeasurement > currentMeasurement {
			counter++
		}

		currentMeasurement = nextMeasurement
	}

	return counter
}

func part2() {
	problem := readFile(inputFilePart1)

	log.Printf("Day 01 part 2: Position of basement: %d", calculateNumberOfThreeMeasurementIncreases(problem))
}

func calculateNumberOfThreeMeasurementIncreases(measurementString string) int {

	measurements := aocstring.SplitToIntArray(measurementString, "\n")

	counter := 0

	a := measurements[0] + measurements[1] + measurements[2]

	for i := 1; i < len(measurements)-2; i++ {
		b := measurements[i] + measurements[i+1] + measurements[i+2]

		if b > a {
			counter++
		}

		a = b
	}

	return counter
}

func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
