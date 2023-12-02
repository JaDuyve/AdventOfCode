package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput(t *testing.T) {
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	expected := 11
	assert.Equalf(t, expected, solvePart1(input), "Sum should be %d.", expected)
}

func Test_SolvePart2_TestInput(t *testing.T) {
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	expected := 6
	assert.Equalf(t, expected, solvePart2(input), "Sum should be %d.", expected)
}
