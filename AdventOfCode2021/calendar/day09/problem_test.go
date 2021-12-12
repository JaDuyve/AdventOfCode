package day09

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	want := 15
	got := solvePart1(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
