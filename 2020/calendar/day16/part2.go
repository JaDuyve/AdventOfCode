package day16

import (
	"AdventOfCode2020/utils/files"
	"regexp"
	"strconv"
)

func Part2() int {
	input := files.ReadFile("calendar/day16/input")
	return solvePart2(input)

}

func solvePart2(puzzleInput string) int {
	fields, ticket, nearbyTickets := parseInput(puzzleInput)
	nearbyTickets = append(nearbyTickets, ticket)
	filterNearbyTickets(fields, &nearbyTickets)

	passedFields := make(map[int]bool)
	history := make(map[string]bool)
	for i := 0; i < len(fields); i++ {
		passedFields[i] = false
	}

	fieldIndexMatch := make([]int, len(fields))
	matchFields(fields, nearbyTickets, &fieldIndexMatch, 0, passedFields, &history)

	return departureProduct(fields, fieldIndexMatch, ticket)
}

func filterNearbyTickets(fields []Field, nearbyTickets *[][]int) {

	for i := len(*nearbyTickets) - 1; i >= 0; i-- {
		if !isValidTicket(fields, (*nearbyTickets)[i]) {
			(*nearbyTickets)[i] = (*nearbyTickets)[len(*nearbyTickets)-1]
			*nearbyTickets = (*nearbyTickets)[:len(*nearbyTickets)-1]
		}
	}
}

func isValidTicket(fields []Field, ticket []int) bool {
	for _, value := range ticket {
		fit := 0

		for _, field := range fields {
			if field.withinRanges(value) {
				fit++
			}
		}

		if fit == 0 {
			return false
		}
	}

	return true
}

// backtracking method
func matchFields(fields []Field, nearbyByTickets [][]int, fieldIndexMap *[]int, valueIndex int, passedField map[int]bool, history *map[string]bool) bool {
	if valueIndex == len(fields) {
		return true
	}

	for i := 0; i < len(fields); i++ {
		if passedField[i] {
			continue
		}

		if isFieldIndexMatch(fields[i], nearbyByTickets, valueIndex, history) {

			(*fieldIndexMap)[i] = valueIndex
			passedField[i] = true

			if matchFields(fields, nearbyByTickets, fieldIndexMap, valueIndex+1, passedField, history) {
				return true
			}

			(*fieldIndexMap)[i] = -1
			passedField[i] = false
		}
	}

	return false
}

func isFieldIndexMatch(field Field, nearbyTickets [][]int, index int, history *map[string]bool) bool {
	if value, exist := (*history)[field.name+strconv.Itoa(index)]; exist {
		return value
	}

	for _, ticket := range nearbyTickets {
		if !field.withinRanges(ticket[index]) {
			(*history)[field.name+strconv.Itoa(index)] = false
			return false
		}
	}

	(*history)[field.name+strconv.Itoa(index)] = true
	return true
}

func departureProduct(fields []Field, fieldIndexMap []int, ticket []int) int {
	departureRegex := regexp.MustCompile("^departure+")
	var product int
	first := true

	for i := 0; i < len(fields); i++ {
		if !departureRegex.MatchString(fields[i].name) {
			continue
		}

		if first {
			first = false
			product = ticket[fieldIndexMap[i]]
			continue
		}

		product *= ticket[fieldIndexMap[i]]

	}

	return product
}
