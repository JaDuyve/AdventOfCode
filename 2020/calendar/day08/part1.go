package day08

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"strings"
)

const (
	acc = "acc"
	jmp = "jmp"
	nop = "nop"
)

type Instruction struct {
	opcode   string
	argument int
	visited  bool
}

func Part1() int {
	input := files.ReadFile("calendar/day08/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	bootCode := buildBootCode(strings.Split(puzzleInput, "\n"))

	accValue, _ := runBootCode(bootCode)
	return accValue
}

func buildBootCode(input []string) *[]Instruction {
	bootCode := make([]Instruction, 0, len(input))

	for _, instruction := range input {
		bootCode = append(bootCode, *createInstruction(instruction))
	}

	return &bootCode
}

func createInstruction(instructionString string) *Instruction {
	return &Instruction{
		opcode:   getOpcode(instructionString),
		argument: getArgument(instructionString),
		visited:  false,
	}
}

func getOpcode(instructionString string) string {
	return instructionString[:3]
}

func getArgument(instructionString string) int {
	return conv.ToInt(instructionString[4:])
}

func runBootCode(bootCode *[]Instruction) (int, int) {
	defer resetBootCode(bootCode)

	accumulator := 0
	instructionPointer := 0

	for instructionPointer != len(*bootCode) {
		instruction := &(*bootCode)[instructionPointer]
		if instruction.visited {
			return accumulator, instructionPointer
		}

		switch instruction.opcode {
		case acc:
			accumulator += instruction.argument
			instructionPointer++
		case jmp:
			instructionPointer += instruction.argument
		case nop:
			instructionPointer++
		}

		instruction.visited = true
	}

	return accumulator, -1
}
