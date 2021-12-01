package day17

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
	"fmt"
	"strings"
)

const frameChar = "."

func Part1() int {
	input := files.ReadFile("calendar/day17/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	cube := enlargeCube(parseInput(puzzleInput))
	cycles := 6

	for i := 0; i < cycles; i++ {
		cube = runCycle(cube)
	}

	return countActiveCubes(&cube)
}

func parseInput(input string) [][][]string {
	rows := strings.Split(input, "\n")
	square := make([][]string, 0, len(rows))

	for _, row := range rows {
		cubeRow := make([]string, 0, 3)

		for _, point := range strings.Split(row, "") {
			cubeRow = append(cubeRow, point)
		}

		square = append(square, cubeRow)
	}

	cube := make([][][]string, 0, 1)

	return append(cube, square)
}

func runCycle(cube [][][]string) [][][]string {
	cube = enlargeCube(cube)
	newCube := slices.CopyString3Dim(cube)

	for z := 1; z < len(cube)-1; z++ {
		for y := 1; y < len(cube[z])-1; y++ {
			for x := 1; x < len(cube[z][y])-1; x++ {
				activeNeighbors := countActiveNeighbors(x, y, z, &cube)
				switch cube[z][y][x] {
				case ".":
					if activeNeighbors == 3 {
						newCube[z][y][x] = "#"
					}
				case "#":
					if activeNeighbors != 2 && activeNeighbors != 3 {
						newCube[z][y][x] = "."
					}
				}
			}
		}
	}

	return newCube
}

func enlargeCube(cube [][][]string) [][][]string {
	newCube := make([][][]string, 0, len(cube)+2)
	paddingSquare := generatePaddingSquare(len(cube[0]) + 2)

	newCube = append(newCube, paddingSquare)
	newCube = append(newCube, cube...)
	newCube = append(newCube, paddingSquare)

	for i := 1; i < len(newCube)-1; i++ {
		newCube[i] = enlargeSquare(newCube[i])
	}

	return newCube
}

func enlargeSquare(square [][]string) [][]string {
	newSquare := make([][]string, 0, len(square)+2)
	padding := generatePaddingRow(len(square) + 2)

	newSquare = append(newSquare, padding)
	newSquare = append(newSquare, square...)
	newSquare = append(newSquare, padding)

	for i := 1; i < len(newSquare)-1; i++ {
		newSquare[i] = enlargeRow(newSquare[i])
	}

	return newSquare
}

func enlargeRow(row []string) []string {
	newRow := make([]string, 0, len(row)+2)
	newRow = append(newRow, frameChar)
	newRow = append(newRow, row...)
	return append(newRow, frameChar)
}

func generatePaddingSquare(width int) [][]string {
	paddingSquare := make([][]string, 0, width)

	for i := 0; i < width; i++ {
		paddingSquare = append(paddingSquare, generatePaddingRow(width))
	}

	return paddingSquare
}

func generatePaddingRow(length int) []string {
	padding := make([]string, 0, length)

	for i := 0; i < length; i++ {
		padding = append(padding, frameChar)
	}

	return padding
}

func countActiveNeighbors(x int, y int, z int, cube *[][][]string) int {
	activeNeighbors := 0

	for _, direction := range directions3D {
		if (*cube)[z+direction[2]][y+direction[1]][x+direction[0]] == "#" {
			activeNeighbors++
		}
	}

	return activeNeighbors
}

func countActiveCubes(cube *[][][]string) int {
	activeCubes := 0

	for z := 1; z < len(*cube)-1; z++ {
		for y := 1; y < len((*cube)[z])-1; y++ {
			for x := 1; x < len((*cube)[z][y])-1; x++ {
				if (*cube)[z][y][x] == "#" {
					activeCubes++
				}
			}
		}
	}

	return activeCubes
}

func printCube(cube *[][][]string) {
	for z := 0; z < len(*cube); z++ {
		fmt.Printf("z = %d\n", z)

		for y := 0; y < len((*cube)[z]); y++ {
			fmt.Printf("%v\n", (*cube)[z][y])
		}
	}
}
