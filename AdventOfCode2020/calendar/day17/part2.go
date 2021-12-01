package day17

import (
	"AdventOfCode2020/utils/files"
	"AdventOfCode2020/utils/slices"
)

func Part2() int {
	input := files.ReadFile("calendar/day17/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	tesseract := parseInput4D(puzzleInput)

	cycles := 6

	for i := 0; i < cycles; i++ {
		tesseract = runCycle4D(tesseract)
	}

	return countActiveCubes4D(&tesseract)
}

func parseInput4D(input string) [][][][]string {
	//cube := enlargeCube(parseInput(input))
	//paddingCube := generatePaddingCube(len(cube))
	//tesseract := make([][][][]string, 0, 1)
	return nil
}

func runCycle4D(tesseract [][][][]string) [][][][]string {
	newCube := slices.CopyString4Dim(tesseract)

	for w := 1; w < len(tesseract)-1; w++ {
		for z := 1; z < len(tesseract[w])-1; z++ {
			for y := 1; y < len(tesseract[w][z])-1; y++ {
				for x := 1; x < len(tesseract[w][z][y])-1; x++ {
					activeNeighbors := countActiveNeighbors4D(x, y, z, w, &tesseract)
					switch tesseract[w][z][y][x] {
					case ".":
						if activeNeighbors == 3 {
							newCube[w][z][y][x] = "#"
						}
					case "#":
						if activeNeighbors != 2 && activeNeighbors != 3 {
							newCube[w][z][y][x] = "."
						}
					}
				}
			}
		}
	}
	tesseract = enlargeTesseract(tesseract)

	return newCube
}

func enlargeTesseract(tesseract [][][][]string) [][][][]string {
	newTesseract := make([][][][]string, 0, len(tesseract)+2)
	paddingCube := generatePaddingCube(len(tesseract[0]) + 2)

	newTesseract = append(newTesseract, paddingCube)
	newTesseract = append(newTesseract, tesseract...)
	newTesseract = append(newTesseract, paddingCube)

	for i := 1; i < len(newTesseract)-1; i++ {
		newTesseract[i] = enlargeCube(newTesseract[i])
	}

	return newTesseract
}

func generatePaddingCube(width int) [][][]string {
	paddingCube := make([][][]string, 0, width)

	for i := 0; i < width; i++ {
		paddingCube = append(paddingCube, generatePaddingSquare(width))
	}

	return paddingCube
}

func countActiveCubes4D(tesseract *[][][][]string) int {
	activeCubes := 0

	for w := 1; w < len(*tesseract)-1; w++ {
		activeCubes += countActiveCubes(&(*tesseract)[w])
	}

	return activeCubes
}

func countActiveNeighbors4D(x, y, z, w int, tesseract *[][][][]string) int {
	activeNeighbors := 0

	for _, direction := range directions4D {
		if (*tesseract)[w+direction[3]][z+direction[2]][y+direction[1]][x+direction[0]] == "#" {
			activeNeighbors++
		}
	}

	return activeNeighbors
}
