package day08

import (
	. "AdventOfCode2021/calendar/day08/PatternMap"
	"AdventOfCode2021/util/aocstring"
	"log"
	"math"
	"regexp"
	"strings"
)

const (
	inputFile = "calendar/day08/input"
)

var expr = regexp.MustCompile("\\S+")

var numberDisp = [][]bool{
	{true, true, true, true, true, false, true},     // 0
	{false, true, true, false, false, false, false}, // 1
	{true, true, false, true, true, true, false},    // 2
	{true, true, true, true, false, true, false},    // 3
	{false, true, true, false, false, true, true},   // 4
	{true, false, true, true, false, true, true},    // 5
	{true, false, true, true, true, true, true},     // 6
	{true, true, true, false, false, false, false},  // 7
	{true, true, true, true, true, true, true},      // 8
	{true, true, true, true, false, true, true},     // 9
}

var numbTwoThreeFive = []bool{
	true, false, false, true, false, true, false,
}
var numbZeroSixNine = []bool{
	true, false, true, true, false, false, true,
}

func Run() {
	problem := aocstring.ReadProblemString(inputFile)

	part1(problem)
	part2(problem)
}

func part1(problem string) {
	log.Printf("Day 08 part 1: Seven Segment Search : %d", solvePart1(problem))
}

func part2(problem string) {
	log.Printf("Day 08 part 2: Seven segment search : %d", solvePart2(problem))
}

func solvePart1(problem string) int {
	lines := strings.Split(problem, "\n")

	freq := make([]int, 8, 8)

	for _, l := range lines {
		countUniqueNumbers(l, &freq)
	}

	counter := freq[3]
	counter += freq[4]
	counter += freq[7]
	counter += freq[2]

	return counter
}

func countUniqueNumbers(l string, freq *[]int) {
	a := strings.Split(l, "|")
	patterns := expr.FindAllString(a[1], -1)

	for _, pattern := range patterns {
		(*freq)[len(pattern)]++
	}
}

func solvePart2(problem string) int {
	lines := strings.Split(problem, "\n")
	counter := 0

	for _, l := range lines {
		s := strings.Split(l, " | ")
		counter += calculatePatternValue(s[0], s[1])
	}

	return counter
}

func calculatePatternValue(patternStr string, digitStr string) int {
	display := make([]byte, 7, 7)
	patternDigits := strings.Split(patternStr, " ")
	letters := make(map[byte]struct{})

	success := sudokuMethod(display, patternDigits, letters, 0)
	if !success {
		log.Fatalf("un able to find display mapping for pattern %s", patternStr)
	}

	pattern := mapDisplayToPattern(display, patternDigits)

	sum := 0

	digits := strings.Split(digitStr, " ")
	for i, digit := range aocstring.Reverse(digits) {
		if i == 0 {
			sum += pattern.GetValue(digit)
		} else {

			value := pattern.GetValue(digit)
			sum += int(math.Pow10(i)) * value
		}
	}

	return sum
}

func mapDisplayToPattern(display []byte, numbers []string) Pattern {
	pattern := New()

	for _, number := range numbers {
		if check(display, number, numberDisp[0], true) {
			pattern.Add(number, 0)
		} else if check(display, number, numberDisp[1], true) {
			pattern.Add(number, 1)
		} else if check(display, number, numberDisp[2], true) {
			pattern.Add(number, 2)
		} else if check(display, number, numberDisp[3], true) {
			pattern.Add(number, 3)
		} else if check(display, number, numberDisp[4], true) {
			pattern.Add(number, 4)
		} else if check(display, number, numberDisp[5], true) {
			pattern.Add(number, 5)
		} else if check(display, number, numberDisp[6], true) {
			pattern.Add(number, 6)
		} else if check(display, number, numberDisp[7], true) {
			pattern.Add(number, 7)
		} else if check(display, number, numberDisp[8], true) {
			pattern.Add(number, 8)
		} else if check(display, number, numberDisp[9], true) {
			pattern.Add(number, 9)
		}
	}

	return pattern
}

func sudokuMethod(display []byte, numbers []string, letters map[byte]struct{}, p int) bool {
	if p == 7 {
		return true
	}

	var i byte
	for i = 'a'; i <= 'g'; i++ {
		if _, ok := letters[i]; ok {
			continue
		}

		display[p] = i

		for x := p + 1; x < len(display); x++ {
			display[x] = 0
		}

		if checkNumbers(display, numbers) {
			letters[i] = struct{}{}
			ok := sudokuMethod(display, numbers, letters, p+1)

			if ok {
				return true
			}

			delete(letters, i)
		}
	}

	return false
}

func checkNumbers(display []byte, numbers []string) bool {

	for _, number := range numbers {
		if !checkNumber(display, number) {
			return false
		}
	}

	return true
}

func checkNumber(display []byte, number string) bool {
	switch len(number) {
	case 2:
		return check(display, number, numberDisp[1], false)
	case 3:
		return check(display, number, numberDisp[7], false)
	case 4:
		return check(display, number, numberDisp[4], false)
	case 7:
		return check(display, number, numberDisp[8], false)
	case 5:
		return check(display, number, numbTwoThreeFive, false)
	case 6:
		return check(display, number, numbZeroSixNine, false)
	}

	log.Fatalf("number has invalid length '%s'", number)
	return false
}

func check(display []byte, number string, symbol []bool, strict bool) bool {
	for i, b := range symbol {
		if display[i] == 0 {
			continue
		}

		contains := aocstring.ContainsChar(number, display[i])
		if b && !contains {
			return false
		} else if strict && !b && contains {
			return false
		}
	}

	return true
}
