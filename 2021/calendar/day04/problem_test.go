package day04

import (
	"AdventOfCode2021/util/aocstring"
	"testing"
)

var input = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestSolvePart1_ExampleProblem(t *testing.T) {
	want := 4512

	got := solvePart1(input)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}

func TestSolvePart1_ColumnWin_ExampleProblem(t *testing.T) {
	want := 4512
	columnWinInput := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

 7  5 20 19  4
10 16 15  9 24
18  8 23 26 17
22 11 13  6 21
 2  0 12  3 14`

	got := solvePart1(columnWinInput)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}

func TestSolvePart2_ExampleProblem(t *testing.T) {
	want := 1924

	got := solvePart2(input)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}

func TestSolvePart1(t *testing.T) {
	want := 89001
	problem := aocstring.ReadProblemString("input")

	got := solvePart1(problem)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	want := 7296
	problem := aocstring.ReadProblemString("input")

	got := solvePart2(problem)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}
