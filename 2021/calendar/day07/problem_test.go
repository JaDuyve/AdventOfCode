package day07

import "testing"

func TestSolvePart1(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	want := 37
	got := solvePart1(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	want := 168
	got := solvePart2(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
