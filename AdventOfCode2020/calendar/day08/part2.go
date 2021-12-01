package day08

import (
	"AdventOfCode2020/utils/files"
	"strings"
)

func Part2() int {
	input := files.ReadFile("calendar/day08/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	bootCode := buildBootCode(strings.Split(puzzleInput, "\n"))
	fixBootCode(bootCode)

	accValue, _ := runBootCode(bootCode)
	return accValue
}

func fixBootCode(bootCode *[]Instruction) {

i:
	for i := 0; i < len(*bootCode); i++ {
		instruction := &(*bootCode)[i]
		switch instruction.opcode {
		case jmp:
			instruction.opcode = nop
			if _, p := runBootCode(bootCode); p == -1 {
				break i
			}

			instruction.opcode = jmp
		case nop:
			instruction.opcode = jmp
			if _, p := runBootCode(bootCode); p == -1 {
				break i
			}

			instruction.opcode = nop
		}
	}
}

func resetBootCode(bootCode *[]Instruction) {
	for i := 0; i < len(*bootCode); i++ {
		(*bootCode)[i].visited = false
	}
}
