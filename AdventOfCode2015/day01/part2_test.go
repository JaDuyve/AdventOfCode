package day01

import "testing"

func Test_Problem2(t *testing.T) {
	tests := []struct {
		problem            string
		positionOfBasement int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		got := solveProblem2(test.problem)

		if test.positionOfBasement != got {
			t.Fatalf("Expected %d but got %d for %s", test.positionOfBasement, got, test.problem)
		}
	}
}
