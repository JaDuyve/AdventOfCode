package day06

import "testing"

func TestSolvePart1_ExampleProblem(t *testing.T) {
	input := "3,4,3,1,2"
	days := 80
	want := 5934

	got := solvePart1(input, days)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestSolvePart2_ExampleProblem(t *testing.T) {
	input := "3,4,3,1,2"
	days := 256
	want := 26984457539

	got := solvePart1(input, days)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
