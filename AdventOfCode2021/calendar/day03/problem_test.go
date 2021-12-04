package day03

import (
	"AdventOfCode2021/util/aocstring"
	"testing"
)

var testNumbers = []int{
	0b00100,
	0b11110,
	0b10110,
	0b10111,
	0b10101,
	0b01111,
	0b00111,
	0b11100,
	0b10000,
	0b11001,
	0b00010,
	0b01010,
}

func TestConvertBinaryStrings(t *testing.T) {
	binaryStrings := []string{"10110", "00100", "11110"}
	numbers := []int{22, 4, 30}

	result := convertBinaryStrings(binaryStrings)

	if len(numbers) != len(result) {
		t.Fatalf("not all binarystrings is in result, wanted size '%d' but got '%d'", len(numbers), len(result))
	}

	for i, number := range result {
		if numbers[i] != number {
			t.Fatalf("wrong convert for '%s', got '%d' but wanted '%d'", binaryStrings[i], number, numbers[i])
		}
	}
}

func TestCounter1And0(t *testing.T) {
	tests := []struct {
		bitPos   int
		counter0 int
		counter1 int
	}{
		{4, 5, 7},
	}

	for _, test := range tests {
		c0, c1 := counter0And1(testNumbers, test.bitPos)

		if c0 != test.counter0 {
			t.Fatalf("counter0 wanted '%d' but got '%d'", test.counter0, c0)
		}

		if c1 != test.counter1 {
			t.Fatalf("counter1 wanted '%d' but got '%d'", test.counter1, c1)
		}
	}

}

func TestCalculateGammaRate(t *testing.T) {
	want := 22
	got := CalculateGammaRate(testNumbers, 5)

	if want != got {
		t.Fatalf("Wrong gamma rate, wanted '%d' (binary '%b') got '%d' (binary '%b')", want, want, got, got)
	}
}

func TestCalculateEpsilonRate(t *testing.T) {
	want := 9
	got := CalculateEpsilonRate(testNumbers, 5)

	if want != got {
		t.Fatalf("Wrong gamma rate, wanted '%d' (binary '%b') got '%d' (binary '%b')", want, want, got, got)
	}
}

func TestSolvePart1_TestProblem(t *testing.T) {
	problem := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	want := 198
	got := solvePart1(problem)

	if want != got {
		t.Fatalf("SolvePart1, want '%d', got '%d'", want, got)
	}
}

func TestSolvePart1(t *testing.T) {
	problem := aocstring.ReadProblemString("input")

	want := 2498354
	got := solvePart1(problem)

	if want != got {
		t.Fatalf("SolvePart1, want '%d', got '%d'", want, got)
	}
}

func TestCalculateOxygenGeneratorRating(t *testing.T) {
	want := 23
	got := calculateOxygenGeneratorRating(testNumbers, 5)

	if want != got {
		t.Fatalf("OxygenGeneratorRating wanted '%d' (binary '%05b') but got '%d' (binary '%05b')", want, want, got, got)
	}
}

func TestCalculateCo2ScrubberRating(t *testing.T) {
	want := 10
	got := calculateCo2ScrubberRating(testNumbers, 5)

	if want != got {
		t.Fatalf("OxygenGeneratorRating wanted '%d' (binary '%05b') but got '%d' (binary '%05b')", want, want, got, got)
	}
}

func TestSolvePart2(t *testing.T) {
	problem := aocstring.ReadProblemString("input")

	want := 3277956
	got := solvePart2(problem)

	if want != got {
		t.Fatalf("SolvePart2, want '%d', got '%d'", want, got)
	}
}

func TestSolvePart2_TestProblem(t *testing.T) {
	problem := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	want := 230
	got := solvePart2(problem)

	if want != got {
		t.Fatalf("SolvePart2, want '%d', got '%d'", want, got)
	}
}
