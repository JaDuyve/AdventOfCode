package day01

import "testing"

func Test_Problem1(t *testing.T) {
	tests := []struct {
		problem          string
		destinationFloor int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		got := solveProblem1(test.problem)

		if test.destinationFloor != got {
			t.Fatalf("Expected %d but got %d for %s", test.destinationFloor, got, test.problem)
		}
	}
}
