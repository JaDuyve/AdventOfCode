package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	input := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"

	t.Run("part 1", func(t *testing.T) {
		expected := 24000
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 45000
		actual := part2(input)

		assert.Equal(actual, expected)
	})
}
