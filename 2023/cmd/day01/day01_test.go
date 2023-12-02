package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	asrt := assert.New(t)
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

	t.Run("part 1", func(t *testing.T) {
		expected := 142
		actual := part1(input)

		asrt.Equal(expected, actual)
	})

	input = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	t.Run("part 2", func(t *testing.T) {
		expected := 281
		actual := part2(input)

		asrt.Equal(expected, actual)
	})
}
