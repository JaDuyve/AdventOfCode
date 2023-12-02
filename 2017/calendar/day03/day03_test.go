package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolveProblem1Example1(t *testing.T) {
	input := 1.0

	want := 0
	got := solveProblem1(input)

	assert.Equal(t, want, got)
}

func Test_SolveProblem1Example2(t *testing.T) {
	input := 12.0

	want := 3
	got := solveProblem1(input)

	assert.Equal(t, want, got)
}

func Test_CalculateCubeSize_Give22Want5(t *testing.T) {
	input := 22

	want := 5
	got := calculateCubeSize(input)

	assert.Equal(t, want, got)
}

func Test_CalculateCubeSize_Give12Want5(t *testing.T) {
	input := 12

	want := 5
	got := calculateCubeSize(input)

	assert.Equal(t, want, got)
}

func Test_CalculateDistance_Give23Want2(t *testing.T) {
	input := 23

	want := 2
	got := calculateDistance(input)

	assert.Equal(t, want, got)
}

func Test_CalculateDistance_Give12Want3(t *testing.T) {
	input := 12

	want := 3
	got := calculateDistance(input)

	assert.Equal(t, want, got)
}

func Test_CalculateDistance_Give1024Want31(t *testing.T) {
	input := 1024

	want := 31
	got := calculateDistance(input)

	assert.Equal(t, want, got)
}

func Test_generateOuterCircle(t *testing.T) {
	inputBegin := 9
	inputEnd := 25

	want := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	got := generateOuterCircle(inputBegin, inputEnd)

	assert.Equal(t, want, got)
}
