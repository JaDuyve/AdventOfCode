package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolvePart1_TestInput(t *testing.T) {
	input := "939\n7,13,x,x,59,x,31,19"
	expected := 295
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInput1(t *testing.T) {
	input := "939\n7,13,x,x,59,x,31,19"
	expected := 1_068_781
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput2(t *testing.T) {
	input := "939\n17,x,13,19"
	expected := 3_417
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput3(t *testing.T) {
	input := "939\n67,7,59,61"
	expected := 754_018
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput4(t *testing.T) {
	input := "939\n67,x,7,59,61"
	expected := 779_210
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput5(t *testing.T) {
	input := "939\n67,7,x,59,61"
	expected := 1_261_476
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInput6(t *testing.T) {
	input := "939\n1789,37,47,1889"
	expected := 1_202_161_486
	assert.Equal(t, expected, solvePart2(input))
}
