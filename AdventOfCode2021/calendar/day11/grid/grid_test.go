package grid

import "testing"

func TestApplyStep(t *testing.T) {
	input := `11111
19991
19191
19991
11111`

	grid := New(input)
	grid.ApplyStep()

	want := `34543
40004
50005
40004
34543`

	got := grid.ToString()

	if want != got {
		t.Fatalf("wanted: \n%s \nbut got: \n%s", want, got)
	}
}

func TestApplyStep2(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	grid := New(input)
	grid.ApplyStep()

	want := `6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637`

	got := grid.ToString()
	t.Logf("\n%s", got)

	if want != got {
		t.Fatalf("step 1: wanted: \n%s \nbut got: \n%s", want, got)
	}

	want = `8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848`

	grid.ApplyStep()
	got = grid.ToString()

	if want != got {
		t.Fatalf("step 2: wanted: \n%s \nbut got: \n%s", want, got)
	}
}
