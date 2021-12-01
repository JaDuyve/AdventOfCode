package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1_TestInput(t *testing.T) {
	input := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0"
	expected := 165
	assert.Equal(t, expected, solvePart1(input))
}

func Test_Part2_TestInput(t *testing.T) {
	input := "mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1"
	expected := 208
	assert.Equal(t, expected, solvePart2(input))
}
