package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetValueAtIndexFirstChar(t *testing.T) {
	inputString := "1234"

	want := 1
	got := getValue(0, inputString)

	assert.Equal(t, want, got)
}

func Test_GetValueAtIndexLastChar(t *testing.T) {
	inputString := "1234"

	want := 4
	got := getValue(3, inputString)

	assert.Equal(t, want, got)
}

func Test_GetValueAtIndexCircleChar(t *testing.T) {
	inputString := "1234"

	want := 1
	got := getValue(4, inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem1Example1(t *testing.T) {
	inputString := "1122"

	want := 3
	got := solveProblem1(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem1Example2(t *testing.T) {
	inputString := "1111"

	want := 4
	got := solveProblem1(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem1Example3(t *testing.T) {
	inputString := "1234"

	want := 0
	got := solveProblem1(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem1Example4(t *testing.T) {
	inputString := "91212129"

	want := 9
	got := solveProblem1(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem2Example1(t *testing.T) {
	inputString := "1212"

	want := 6
	got := solveProblem2(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem2Example2(t *testing.T) {
	inputString := "1221"

	want := 0
	got := solveProblem2(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem2Example3(t *testing.T) {
	inputString := "123425"

	want := 4
	got := solveProblem2(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem2Example4(t *testing.T) {
	inputString := "123123"

	want := 12
	got := solveProblem2(inputString)

	assert.Equal(t, want, got)
}

func Test_solveProblem2Example5(t *testing.T) {
	inputString := "12131415"

	want := 4
	got := solveProblem2(inputString)

	assert.Equal(t, want, got)
}
