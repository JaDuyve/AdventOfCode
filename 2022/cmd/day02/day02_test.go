package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	input := "A Y\nB X\nC Z"

	t.Run("part 1", func(t *testing.T) {
		expected := 15
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 12
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
