package day03

import (
	"testing"
)

func TestConvertBinaryStrings(t *testing.T) {
	binaryStrings := []string{"10110", "00100", "11110"}
	numbers := []int64{22, 4, 30}

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

func TestCalculateGammaRate(t *testing.T) {
	numbers := []int64{
		int64(0b00100),
		int64(0b11110),
		int64(0b10110),
		int64(0b10111),
		int64(0b10101),
		int64(0b01111),
		int64(0b00111),
		int64(0b11100),
		int64(0b10000),
		int64(0b11001),
		int64(0b00010),
		int64(0b01010),
	}

	want := 22
	got := CalculateGammaRate(numbers, 5)

	if want != got {
		t.Fatalf("Wrong gamma rate, wanted '%d' (binary '%b') got '%d' (binary '%b')", want, want, got, got)
	}
}

func TestCalculateEpsilonRate(t *testing.T) {
	numbers := []int64{
		int64(0b00100),
		int64(0b11110),
		int64(0b10110),
		int64(0b10111),
		int64(0b10101),
		int64(0b01111),
		int64(0b00111),
		int64(0b11100),
		int64(0b10000),
		int64(0b11001),
		int64(0b00010),
		int64(0b01010),
	}

	want := 9
	got := CalculateEpsilonRate(numbers, 5)

	if want != got {
		t.Fatalf("Wrong gamma rate, wanted '%d' (binary '%b') got '%d' (binary '%b')", want, want, got, got)
	}
}

func TestSolvePart1(t *testing.T) {
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
