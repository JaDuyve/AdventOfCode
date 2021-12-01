package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MinMaxExample1(t *testing.T) {
	inputString := "5\t1\t9\t5"

	wantMin := 1
	wantMax := 9

	gotMin, gotMax := minMax(inputString)

	assert.Equal(t, wantMin, gotMin)
	assert.Equal(t, wantMax, gotMax)
}

func Test_MinMaxExample2(t *testing.T) {
	inputString := "7\t5\t3"

	wantMin := 3
	wantMax := 7

	gotMin, gotMax := minMax(inputString)

	assert.Equal(t, wantMin, gotMin)
	assert.Equal(t, wantMax, gotMax)
}

func Test_evenlyDivisibleValue(t *testing.T) {
	inputString := "5\t9\t2\t8"

	want := 4

	got := getEvenlyDivisibleValue(inputString)

	assert.Equal(t, want, got)
}

func Test_SolveProblem1(t *testing.T) {
	inputString := "5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8"

	want := 18
	got := solveProblem1(inputString)

	assert.Equal(t, want, got)
}
