package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1_TestInput(t *testing.T) {
	input := "F10\nN3\nF7\nR90\nF11"
	expected := 25
	assert.Equal(t, expected, solvePart1(input))
}

func Test_Part2_TestInput(t *testing.T) {
	input := "F10\nN3\nF7\nR90\nF11"
	expected := 286
	assert.Equal(t, expected, solvePart2(input))
}
