package day10

import "testing"

func TestIsChunkValid(t *testing.T) {
	tests := []struct {
		chunk          string
		isValid        bool
		invalidBracket byte
	}{
		{"([])", true, '.'},
		{"{()()()}", true, '.'},
		{"<([{}])>", true, '.'},
		{"[<>({}){}[([])<>]]", true, '.'},
		{"(((((((((())))))))))", true, '.'},
		{"{([(<{}[<>[]}>{[]{[(<()>", false, '}'},
		{"[[<[([]))<([[{}[[()]]]", false, ')'},
		{"[{[{({}]{}}([{[{{{}}([]", false, ']'},
		{"[<(<(<(<{}))><([]([]()", false, ')'},
		{"<{([([[(<>()){}]>(<<{{", false, '>'},
	}

	for _, test := range tests {
		isValid, ch := isChunkValid(test.chunk)

		if isValid != test.isValid {
			t.Fatalf("Wanted '%v' but got '%v' for chunk '%s'", test.isValid, isValid, test.chunk)
		}

		if test.isValid == false && test.invalidBracket != ch {
			t.Fatalf("wanted '%c' but got '%c' for chunk '%s'", test.invalidBracket, ch, test.chunk)
		}
	}
}

func TestFixChunk(t *testing.T) {
	tests := []struct {
		chunk         string
		completionStr string
	}{
		{"[({(<(())[]>[[{[]{<()<>>", "}}]])})]"},
		{"[(()[<>])]({[<{<<[]>>(", ")}>]})"},
		{"(((({<>}<{<{<>}{[]{[]{}", "}}>}>))))"},
		{"{<[[]]>}<{[{[{[]{()[[[]", "]]}}]}]}>"},
		{"<{([{{}}[<[[[<>{}]]]>[]]", "])}>"},
	}

	for _, test := range tests {
		compArr := fixChunk(test.chunk)

		compStr := string(compArr)

		if compStr != test.completionStr {
			t.Fatalf("wanted '%s' but got '%s' for chunk '%s'", test.completionStr, compStr, test.chunk)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	want := 26397
	got := solvePart1(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	want := 288957
	got := solvePart2(input)

	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}

func TestCalculatePointsForChunk(t *testing.T) {
	input := "<{([{{}}[<[[[<>{}]]]>[]]"

	want := 294
	got := calculatePointsForChunk(input)

	if want != got {
		t.Fatalf("wanted '%d' but got '%d'", want, got)
	}
}
