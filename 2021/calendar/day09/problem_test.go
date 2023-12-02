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

func TestCalculateSizeBasin(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	heightMap := parseInput(input)

	tests := []struct {
		y    int
		x    int
		size int
	}{
		{0, 9, 9},
		{0, 0, 3},
		{2, 2, 14},
		{4, 6, 9},
	}

	for _, test := range tests {
		got := calculateSizeBasin(test.y, test.x, &heightMap)

		if test.size != got {
			t.Fatalf("wanted %d but got %d", test.size, got)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	want := 1134
	got := solvePart2(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
