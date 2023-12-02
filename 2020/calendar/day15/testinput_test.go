package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput1(t *testing.T) {
	input := "0,3,6"
	expected := 436
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput2(t *testing.T) {
	input := "1,3,2"
	expected := 1
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput3(t *testing.T) {
	input := "2,1,3"
	expected := 10
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput4(t *testing.T) {
	input := "1,2,3"
	expected := 27
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput5(t *testing.T) {
	input := "2,3,1"
	expected := 78
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput6(t *testing.T) {
	input := "3,2,1"
	expected := 438
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart1_TestInput7(t *testing.T) {
	input := "3,1,2"
	expected := 1836
	assert.Equal(t, expected, solvePart1(input, 2020))
}

func Test_SolvePart2_TestInput1(t *testing.T) {
	input := "0,3,6"
	expected := 175594
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput2(t *testing.T) {
	input := "1,3,2"
	expected := 2578
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput3(t *testing.T) {
	input := "2,1,3"
	expected := 3544142
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput4(t *testing.T) {
	input := "1,2,3"
	expected := 261214
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput5(t *testing.T) {
	input := "2,3,1"
	expected := 6895259
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput6(t *testing.T) {
	input := "3,2,1"
	expected := 18
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput7(t *testing.T) {
	input := "3,1,2"
	expected := 362
	assert.Equal(t, expected, solvePart2(input))
}
