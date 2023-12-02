package main

import "fmt"

func evaluateCode(code int) bool {
	return sixDigitRule(code) &&
		doubleDigitRule(code) &&
		numberIncreaseRule(code)
}

func sixDigitRule(code int) bool {
	return code > 99999 && code < 1000000
}

func doubleDigitRule(code int) bool {

	for code > 0 {
		firstDigit := code % 10
		code = code / 10
		secondDigit := code % 10

		if firstDigit == secondDigit {
			return true
		}
	}

	return false
}

func numberIncreaseRule(code int) bool {

	for code > 0 {
		firstDigit := code % 10
		code = code / 10
		secondDigit := code % 10

		if secondDigit > firstDigit {
			return false
		}
	}

	return true
}

func main() {
	lowerbound := 245318
	upperbound := 765747
	countValidNumbers := 0

	for a := lowerbound; a < upperbound; a++ {
		if evaluateCode(a) {
			countValidNumbers++
		}
	}

	fmt.Printf("Number of valid codes: %d\n", countValidNumbers)
}
