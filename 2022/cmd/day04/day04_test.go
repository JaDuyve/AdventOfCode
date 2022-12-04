package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"

	t.Run("part 1", func(t *testing.T) {
		expected := 2
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 4
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
