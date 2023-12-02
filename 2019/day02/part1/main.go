package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

var intcodeDivider = ','

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input")
	Check(err)

	program := string(dat)

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	numberStrings := strings.FieldsFunc(program, f)

	numbers := make([]int, len(numberStrings))
	for index, str := range numberStrings {
		number, err := strconv.Atoi(str)
		Check(err)
		numbers[index] = number
	}

	numbers[1] = 12
	numbers[2] = 2

	currentPosition := 0
	for currentPosition < len(numbers) && numbers[currentPosition] != 99 {
		if numbers[currentPosition] == 1 {
			numbers[numbers[currentPosition+3]] = numbers[numbers[currentPosition+1]] + numbers[numbers[currentPosition+2]]
		} else if numbers[currentPosition] == 2 {
			numbers[numbers[currentPosition+3]] = numbers[numbers[currentPosition+1]] * numbers[numbers[currentPosition+2]]
		}
		currentPosition += 4
	}

	fmt.Printf("numbers[0]: %d\n", numbers[0])
}
