package day13

import "testing"

const input = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestNewPaper(t *testing.T) {
	p := NewPaper(input)

	gridStr := p.ToString()
	const expectedGridStr = `...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........`

	if gridStr != expectedGridStr {
		t.Fatalf("wanted: \n%s\n but got: \n%s", expectedGridStr, gridStr)
	}
}

func TestFold(t *testing.T) {
	const expectedGridStr = `#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........`

	p := NewPaper(input)
	p.Fold()
	gridStr := p.ToString()

	if gridStr != expectedGridStr {
		t.Fatalf("wanted: \n%s\nbut got:\n%s", expectedGridStr, gridStr)
	}
}

func TestFold2(t *testing.T) {
	const expectedGridStr = `#####
#...#
#...#
#...#
#####
.....
.....`

	p := NewPaper(input)
	p.Fold()
	p.Fold()
	gridStr := p.ToString()

	if gridStr != expectedGridStr {
		t.Fatalf("wanted: \n%s\nbut got:\n%s", expectedGridStr, gridStr)
	}
}
