package day14

import "testing"

func TestApplySteps1(t *testing.T) {
	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	poly := New(input, 10)

	poly.ApplySteps()

	want := 1588
	got := poly.GetMinMaxDifference()

	if want != got {
		t.Fatalf("\nwant %d but got %d", want, got)
	}

}
