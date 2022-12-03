package day13

import (
	"AdventOfCode2021/util/aocmath"
	"fmt"
	"log"
	"strings"
)

type Paper struct {
	Grid         [][]byte
	Instructions []Instruction
	Pos          int
	maxY         int
	maxX         int
}

type Instruction struct {
	Axis     byte
	Position int
}

func NewPaper(input string) *Paper {
	s := strings.Split(input, "\n\n")
	gridCoordinates := strings.Split(s[0], "\n")
	instructionsStr := s[1]

	xs := make([]int, len(gridCoordinates), len(gridCoordinates))
	ys := make([]int, len(gridCoordinates), len(gridCoordinates))
	for i, gridCo := range gridCoordinates {
		y, x := 0, 0

		_, err := fmt.Sscanf(gridCo, "%d,%d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}

		xs[i] = x
		ys[i] = y
	}

	maxX := aocmath.MaxOfArray(xs)
	maxY := aocmath.MaxOfArray(ys)

	grid := createGrid(maxY+1, maxX+1)

	for i := 0; i < len(ys); i++ {
		grid[ys[i]][xs[i]] = '#'
	}

	return &Paper{
		grid,
		ParseInstructions(instructionsStr),
		0,
		maxY,
		maxX,
	}
}

func createGrid(sizeY, sizeX int) [][]byte {
	grid := make([][]byte, sizeY)
	for y := 0; y < sizeY; y++ {
		grid[y] = make([]byte, sizeX)
		for x := 0; x < sizeX; x++ {
			grid[y][x] = '.'
		}
	}

	return grid
}

func ParseInstructions(input string) []Instruction {
	instructionStr := strings.Split(input, "\n")

	instructions := make([]Instruction, len(instructionStr), len(instructionStr))
	for i, s := range instructionStr {
		instructions[i] = NewInstruction(s)
	}

	return instructions
}

func NewInstruction(input string) Instruction {
	var axis byte
	var pos int
	_, err := fmt.Sscanf(input, "fold along %c=%d", &axis, &pos)
	if err != nil {
		log.Fatal(err)
	}

	return Instruction{
		axis,
		pos,
	}
}

func (p Paper) ToString() string {
	builder := strings.Builder{}

	for y := 0; y < len(p.Grid); y++ {
		if builder.Len() != 0 {
			builder.WriteString("\n")
		}
		for x := 0; x < len(p.Grid[y]); x++ {
			builder.WriteByte(p.Grid[y][x])
		}
	}

	return builder.String()
}

func (p *Paper) Fold() {
	instruction := p.Instructions[p.Pos]

	if instruction.Axis == 'y' {
		p.foldY(instruction.Position)
	} else if instruction.Axis == 'x' {
		p.foldX(instruction.Position)
	}

	p.Pos++
}

func (p *Paper) foldY(pos int) {
	newGrid := createGrid(pos, p.maxX+1)

	for y, yy := 0, len(p.Grid)-1; y < pos && yy > pos; y, yy = y+1, yy-1 {
		for x := 0; x < len(p.Grid[y]); x++ {
			if p.Grid[yy][x] == '#' {
				newGrid[y][x] = '#'
			} else if p.Grid[y][x] == '#' {
				newGrid[y][x] = '#'
			}
		}
	}

	p.maxY = pos - 1
	p.Grid = newGrid
}

func (p *Paper) foldX(pos int) {
	newGrid := createGrid(p.maxY+1, pos)

	for y := 0; y < len(p.Grid); y++ {
		for x, xx := 0, len(p.Grid[y])-1; x < pos && xx > pos; x, xx = x+1, xx-1 {
			if p.Grid[y][xx] == '#' {
				newGrid[y][x] = '#'
			} else if p.Grid[y][x] == '#' {
				newGrid[y][x] = '#'
			}
		}
	}

	p.maxX = pos - 1
	p.Grid = newGrid
}

func (p Paper) CountDots() int {
	counter := 0

	for y := 0; y < len(p.Grid); y++ {
		for x := 0; x < len(p.Grid[y]); x++ {
			if p.Grid[y][x] == '#' {
				counter++
			}
		}
	}

	return counter
}
