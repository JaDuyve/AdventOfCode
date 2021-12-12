package day05

import "testing"

func TestCalcIntersectionPoint(t *testing.T) {
	expectedX := 7
	expectedY := 4

	line1 := line{
		point{7, 0},
		point{7, 4},
	}
	line2 := line{
		point{9, 4},
		point{3, 4},
	}

	intType, intPoint := calcIntersectionPoint(line1, line2)

	if intType != DO_INTERSECT {
		t.Fatalf("expected '%d' but got '%d'", DO_INTERSECT, intType)
	}

	if intPoint.x != expectedX || intPoint.y != expectedY {
		t.Fatalf("expected point: [%d,%d] but got [%d,%d]", expectedX, expectedY, intPoint.x, intPoint.y)
	}
}

func TestCalcIntersectionPoint_ColLinear(t *testing.T) {
	//expectedX := 7
	//expectedY := 4

	line1 := line{
		point{0, 9},
		point{5, 9},
	}
	line2 := line{
		point{0, 9},
		point{2, 9},
	}

	intType, _ := calcIntersectionPoint(line1, line2)

	if intType != COLLINEAR {
		t.Fatalf("expected '%d' but got '%d'", DO_INTERSECT, intType)
	}

	//if intPoint.x != expectedX || intPoint.y != expectedY {
	//	t.Fatalf("expected point: [%d,%d] but got [%d,%d]", expectedX, expectedY, intPoint.x, intPoint.y)
	//}
}

func TestCalcIntersectionPoint2(t *testing.T) {

	line1 := line{
		point{2, 2},
		point{2, 1},
	}
	line2 := line{
		point{9, 4},
		point{3, 4},
	}

	intType, _ := calcIntersectionPoint(line1, line2)

	if intType != DONT_INTERSECT {
		t.Fatalf("expected '%d' but got '%d'", DO_INTERSECT, intType)
	}
}

func TestSolvePart1_ExampleProblem(t *testing.T) {
	problem := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	expectedIntersections := 5

	got := solvePart1(problem)

	if expectedIntersections != got {
		t.Fatalf("expected '%d' intersections but got '%d'", expectedIntersections, got)
	}
}

func TestSolvePart1Second_ExampleProblem(t *testing.T) {
	problem := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	expectedIntersections := 5

	got := SolvePart1Second(problem)

	if expectedIntersections != got {
		t.Fatalf("expected '%d' intersections but got '%d'", expectedIntersections, got)
	}
}
