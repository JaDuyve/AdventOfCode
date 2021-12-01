package day14

import (
	"AdventOfCode2020/utils/files"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day14/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	initializeProgram := strings.Split(puzzleInput, "\n")
	memory := make(map[int]int)
	var mask string
	maskRegex := regexp.MustCompile("^mask = (\\w{36})$")

	for _, operation := range initializeProgram {
		if maskRegex.MatchString(operation) {
			mask = maskRegex.FindStringSubmatch(operation)[1]
			continue
		}

		encryptedIndex, value := getIndexValuePair(operation)
		indexes := generateIndexes(mask, encryptedIndex)

		for _, index := range indexes {
			memory[index] = value
		}
	}

	return calculateSum(memory)
}

func generateIndexes(mask string, index int) []int {
	numberOfX := strings.Count(mask, "X")
	indexes := make([]int, 0, int(math.Pow(2, float64(numberOfX))))

	xNumber := 0
	indexString := []rune(fmt.Sprintf("%036b", index))
	xIndexes := make([]int, 0, numberOfX)

	for i := 0; i < 36; i++ {
		if mask[i] == 'X' {
			indexString[i] = 'X'
			xIndexes = append(xIndexes, i)
		} else if mask[i] == '1' {
			indexString[i] = '1'
		}
	}

	for i := 0; float64(i) < math.Pow(2, float64(numberOfX)); i++ {
		xNumberBinary := []rune(fmt.Sprintf("%0"+strconv.Itoa(numberOfX)+"b", xNumber))
		value := indexString

		for j := 0; j < numberOfX; j++ {
			value[xIndexes[j]] = xNumberBinary[j]
		}

		ii, _ := strconv.ParseInt(string(value), 2, 64)
		indexes = append(indexes, int(ii))
		xNumber++
	}

	return indexes
}
