package day08

import (
	"strings"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

	want := 26
	got := solvePart1(input)

	if want != got {
		t.Fatalf("Wanted %d but got %d", want, got)
	}
}

func TestSudokuMethod(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"

	display := make([]byte, 7, 7)
	numbers := strings.Split(input, " ")
	let := make(map[byte]struct{})

	ok := sudokuMethod(display, numbers, let, 0)
	if !ok {
		t.Fatal("no solution found")
	}

	solution := []byte{'d', 'a', 'b', 'c', 'g', 'f', 'e'}
	for i := range display {
		if display[i] != solution[i] {
			t.Fatalf("expected display node '%c' but got '%c'", solution[i], display[i])
		}
	}
}

func TestMapDisplayToPattern(t *testing.T) {
	linePattern := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	display := []byte{'d', 'a', 'b', 'c', 'g', 'f', 'e'}
	numbers := strings.Split(linePattern, " ")

	patternMap := mapDisplayToPattern(display, numbers)

	expectedValues := map[string]int{
		"acedgfb": 8,
		"cdfbe":   5,
		"gcdfa":   2,
		"fbcad":   3,
		"dab":     7,
		"cefabd":  9,
		"cdfgeb":  6,
		"eafb":    4,
		"cagedb":  0,
		"ab":      1,
	}

	for expectedKey, expectedValue := range expectedValues {
		got := patternMap.GetValue(expectedKey)

		if got != expectedValue {
			t.Fatalf("Wanted %d but got %d for key %s", expectedValue, got, expectedKey)
		}
	}
}

func TestCalculatePatternValue(t *testing.T) {
	patternStr := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	digitStr := "cdfeb fcadb cdfeb cdbaf"

	want := 5353
	got := calculatePatternValue(patternStr, digitStr)

	if want != got {
		t.Fatalf("Wanted %d but got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

	want := 61229
	got := solvePart2(input)

	if want != got {
		t.Fatalf("Wanted %d but got %d", want, got)
	}
}
