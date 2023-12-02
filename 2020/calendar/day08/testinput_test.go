package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1_TestInput(t *testing.T) {
	input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	expected := 5
	assert.Equal(t, expected, solvePart1(input))
}

func Test_Part2_TestInput(t *testing.T) {
	input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	expected := 8
	assert.Equal(t, expected, solvePart2(input))
}
