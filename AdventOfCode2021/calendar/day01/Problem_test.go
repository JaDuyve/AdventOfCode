package day01

import "testing"

func Test_CalculateNumberOfIncreases(t *testing.T) {
	want := 7
	input := `199
200
208
210
200
207
240
269
260
263`

	got := calculateNumberOfIncreases(input)

	if want != got {
		t.Fatalf("Want '%d' but got '%d'", want, got)
	}
}

func Test_calculateNumberOfThreeMeasurementIncreases(t *testing.T) {
	want := 5
	input := `199
200
208
210
200
207
240
269
260
263`

	got := calculateNumberOfThreeMeasurementIncreases(input)

	if want != got {
		t.Fatalf("Want '%d' but got '%d'", want, got)
	}
}
