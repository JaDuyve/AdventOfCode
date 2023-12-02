package day11

import (
	"AdventOfCode2021/calendar/day11/grid"
	"testing"
)

func TestApplySteps10Iterations(t *testing.T) {
	n := 10
	want := 204

	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	g := grid.New(input)

	got := applySteps(g, n)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestApplySteps100Iterations(t *testing.T) {
	n := 100
	want := 1656

	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	g := grid.New(input)

	got := applySteps(g, n)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {

	want := 195

	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	got := solvePart2(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
