package day02

import "testing"

func Test_SolvePart1(t *testing.T) {
	want := 150
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	got := solvePart1(input)

	if want != got {
		t.Fatalf("Want '%d' but got '%d'", want, got)
	}
}
