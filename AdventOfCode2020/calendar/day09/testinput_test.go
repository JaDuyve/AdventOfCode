package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Part1_TestInput(t *testing.T) {
	input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
	expected := 127
	assert.Equal(t, expected, solvePart1(input, 5))
}

func Test_Part2_TestInput(t *testing.T) {
	input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
	expected := 62
	assert.Equal(t, expected, solvePart2(input, 5))
}
