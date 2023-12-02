package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1(t *testing.T) {
	input := "..##.......\n" +
		"#...#...#..\n" +
		".#....#..#.\n" +
		"..#.#...#.#\n" +
		".#...##..#.\n" +
		"..#.##.....\n" +
		".#.#.#....#\n" +
		".#........#\n" +
		"#.##...#...\n" +
		"#...##....#\n" +
		".#..#...#.#"
	assert.Equal(t, 7, solvePart1(input), "Answer to the puzzle should be 7.")
}
