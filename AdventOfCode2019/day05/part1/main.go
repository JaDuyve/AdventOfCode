package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var expectedSolution = 19690720

var ADD_OPCODE = 1
var MULTIPLY_OPCODE = 2
var READ_INPUT_OPCODE = 3
var WRITE_OUTPUT_OPCODE = 4
var JUMP_IF_TRUE_OPCODE = 5
var JUMP_IF_FALSE_OPCODE = 6
var LESS_THAN_OPCODE = 7
var EQUALS_OPCODE = 8
var END_PROGRAM_OPCODE = 99

type Parameters struct {
	opcode          int
	firstParameter  int
	secondParameter int
	saveLocation    int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func runProgram(memory []int) {

	currentPosition := 0

	for isEndProgram(currentPosition, memory) {
		params := getParameters(currentPosition, memory)

		if params.opcode == READ_INPUT_OPCODE {
			inputValue(params, memory)
			currentPosition += 2
		} else if params.opcode == WRITE_OUTPUT_OPCODE {
			outputValue(params)
			currentPosition += 2
		} else if params.opcode == ADD_OPCODE {
			add(params, currentPosition, memory)
			currentPosition += 4
		} else if params.opcode == MULTIPLY_OPCODE {
			multiply(params, currentPosition, memory)
			currentPosition += 4
		} else if params.opcode == JUMP_IF_TRUE_OPCODE {
			res := jumpIfTrue(params)
			if res != -1 {
				currentPosition = res
			} else {
				currentPosition += 3
			}
		} else if params.opcode == JUMP_IF_FALSE_OPCODE {
			res := jumpIfFalse(params)
			if res != -1 {
				currentPosition = res
			} else {
				currentPosition += 3
			}
		} else if params.opcode == LESS_THAN_OPCODE {
			lessThan(params, memory)
			currentPosition += 4
		} else if params.opcode == EQUALS_OPCODE {
			equals(params, memory)
			currentPosition += 4
		}
	}
}

func isEndProgram(currentPosition int, program []int) bool {
	return currentPosition < len(program)-1 &&
		immediateModeRead(currentPosition, program) != END_PROGRAM_OPCODE
}

func add(params Parameters, location int, program []int) {

	result := params.firstParameter + params.secondParameter

	writeValue(result, params.saveLocation, program)
}

func multiply(params Parameters, location int, program []int) {

	result := params.firstParameter * params.secondParameter

	writeValue(result, params.saveLocation, program)
}

func inputValue(params Parameters, program []int) {
	var inputValue int

	fmt.Print("Give input value: ")
	n, err := fmt.Scan(&inputValue)

	for err != nil || n != 1 {
		fmt.Print("Give input value: ")
		n, err = fmt.Scan(&inputValue)
	}

	writeValue(inputValue, params.saveLocation, program)
}

func outputValue(params Parameters) {

	fmt.Printf("Diagnostic code: [%d]\n", params.firstParameter)
}

func jumpIfTrue(params Parameters) int {
	if params.firstParameter != 0 {
		return params.secondParameter
	}

	return -1
}

func jumpIfFalse(params Parameters) int {
	if params.firstParameter == 0 {
		return params.secondParameter
	}

	return -1
}

func lessThan(params Parameters, program []int) {
	if params.firstParameter < params.secondParameter {
		writeValue(1, params.saveLocation, program)
	} else {
		writeValue(0, params.saveLocation, program)
	}
}

func equals(params Parameters, program []int) {
	if params.firstParameter == params.secondParameter {
		writeValue(1, params.saveLocation, program)
	} else {
		writeValue(0, params.saveLocation, program)
	}
}

func getParameters(location int, program []int) Parameters {
	var params Parameters

	if immediateModeRead(location, program) > 2 {
		params = initCalcSpecificsForCrypticOpcode(location, program)
	} else {
		params = getParametersNormalOpcode(location, program)
	}

	if params.opcode != 4 && params.opcode != 5 && params.opcode != 6 {
		params.saveLocation = immediateModeRead(location+3, program)
	}

	return params
}

func getParametersNormalOpcode(location int, program []int) Parameters {
	var st Parameters

	st.opcode = immediateModeRead(location, program)
	st.firstParameter = positionModeRead(location+1, program)
	st.secondParameter = positionModeRead(location+2, program)

	return st
}

func initCalcSpecificsForCrypticOpcode(location int, program []int) Parameters {
	var st Parameters

	crypticOpcode := immediateModeRead(location, program)

	// get opcode
	st.opcode = crypticOpcode % 100
	crypticOpcode /= 100

	// get first number
	typ := crypticOpcode % 10
	st.firstParameter = readValue(typ, location+1, program)
	crypticOpcode /= 10

	if st.opcode != 4 && st.opcode != 3 {
		// get second number
		typ = crypticOpcode % 10
		st.secondParameter = readValue(typ, location+2, program)
	}

	return st
}

func writeValue(value int, location int, program []int) {
	program[location] = value
}

func readValue(typ int, location int, program []int) int {
	var value int

	if typ == 0 {
		value = positionModeRead(location, program)
	} else if typ == 1 {
		value = immediateModeRead(location, program)
	}

	return value
}

func positionModeRead(location int, program []int) int {
	return immediateModeRead(program[location], program)
}

func immediateModeRead(location int, program []int) int {
	return program[location]
}

func main() {
	dat, err := ioutil.ReadFile("day05/part1/input")
	check(err)

	program := string(dat)

	var splits = strings.Split(program, ",")

	numbers := make([]int, len(splits))
	for index, str := range splits {
		number, err := strconv.Atoi(str)
		check(err)
		numbers[index] = number
	}

	runProgram(numbers)

}
