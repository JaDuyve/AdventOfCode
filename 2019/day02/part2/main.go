package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

var expectedSolution = 19690720

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func runProgram(memory []int, noun int, verb int) int {

	memory[1] = noun
	memory[2] = verb

	currentPosition := 0
	for currentPosition < len(memory)-1 && memory[currentPosition] != 99 {
		if memory[currentPosition+3] > len(memory) || memory[currentPosition+1] > len(memory) || memory[currentPosition+2] > len(memory) {
			return 0
		}
		if memory[currentPosition] == 1 {
			memory[memory[currentPosition+3]] = memory[memory[currentPosition+1]] + memory[memory[currentPosition+2]]
		} else if memory[currentPosition] == 2 {
			memory[memory[currentPosition+3]] = memory[memory[currentPosition+1]] * memory[memory[currentPosition+2]]
		}
		currentPosition += 4

	}

	return memory[0]
}

func main() {
	dat, err := ioutil.ReadFile("input")
	check(err)

	program := string(dat)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	numberStrings := strings.FieldsFunc(program, f)

	numbers := make([]int, len(numberStrings))
	for index, str := range numberStrings {
		number, err := strconv.Atoi(str)
		check(err)
		numbers[index] = number
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			numbersCopy := make([]int, len(numbers))
			copy(numbersCopy, numbers)
			sol := runProgram(numbersCopy, noun, verb)
			if expectedSolution == sol {
				fmt.Printf("Noun [%d] and verb [%d] produces the output [%d]\n", noun, verb, expectedSolution)
				return
			}
			// fmt.Printf("Noun [%d] and verb [%d] produces the output [%d]\n", noun, verb, sol)
		}
	}

	fmt.Println(numbers)

}
