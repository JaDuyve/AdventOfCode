package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	input := "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"

	t.Run("part 1", func(t *testing.T) {
		expected := "CMZ"
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "MCD"
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}

func TestTransposeString(t *testing.T) {
	input := "    [D]    \n[N] [C]    \n[Z] [M] [P]"
	expected := "[[ \nZN \n]] \n   \n[[[\nMCD\n]]]\n   \n[  \nP  \n]  "

	t.Logf("%s\n", input)
	result := TransposeString(input)
	t.Logf("%s\n", result)

	if expected != result {
		t.Fail()
	}
}

func TestParseState(t *testing.T) {
	input := "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 "

	ParseState(input)
}
