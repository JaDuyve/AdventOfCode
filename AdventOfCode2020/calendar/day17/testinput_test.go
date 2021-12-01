package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput(t *testing.T) {
	input := ".#.\n..#\n###"
	expected := 112
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInput(t *testing.T) {
	input := ".#.\n..#\n###"
	expected := 848
	assert.Equal(t, expected, solvePart2(input))
}
