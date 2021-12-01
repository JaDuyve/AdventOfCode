package day16

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"regexp"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r *Range) withinRange(value int) bool {
	return value >= r.start && value <= r.end
}

type Field struct {
	name   string
	ranges []Range
}

func (f *Field) withinRanges(value int) bool {
	fit := 0

	for _, r := range f.ranges {
		if r.withinRange(value) {
			fit++
		}
	}

	return fit > 0
}

func Part1() int {
	input := files.ReadFile("calendar/day16/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {

	fields, _, nearbyTickets := parseInput(puzzleInput)

	return calculateScanningErrorRate(nearbyTickets, fields)
}

func parseInput(input string) ([]Field, []int, [][]int) {
	splitInput := strings.Split(input, "\n\n")

	textFields := splitInput[0]
	textTicket := splitInput[1]
	textNearbyTickets := splitInput[2]

	fields := extractFields(textFields)
	ticket := extractTicket(strings.Split(textTicket, "\n")[1])
	nearbyTickets := extractNearbyTickets(textNearbyTickets)

	return fields, ticket, nearbyTickets
}

func extractFields(textFields string) []Field {
	splitFields := strings.Split(textFields, "\n")
	fields := make([]Field, 0, len(textFields))

	fieldRegex := regexp.MustCompile("(?m)^([\\w\\s]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")

	for _, fieldString := range splitFields {
		fieldValues := fieldRegex.FindAllStringSubmatch(fieldString, -1)[0][1:]

		field := new(Field)
		field.name = fieldValues[0]

		for i := 1; i < len(fieldValues); i += 2 {
			field.ranges = append(field.ranges, Range{
				start: conv.ToInt(fieldValues[i]),
				end:   conv.ToInt(fieldValues[i+1]),
			})
		}

		fields = append(fields, *field)
	}

	return fields
}

func extractTicket(ticketString string) []int {
	return conv.ToIntSlice(strings.Split(ticketString, ","))
}

func extractNearbyTickets(nearbyTicketString string) [][]int {
	nearbyTickets := make([][]int, 0)

	for _, ticket := range strings.Split(nearbyTicketString, "\n")[1:] {
		nearbyTickets = append(nearbyTickets, extractTicket(ticket))
	}

	return nearbyTickets
}

func calculateScanningErrorRate(nearbyTickets [][]int, fields []Field) int {
	sum := 0

	for i := 0; i < len(nearbyTickets); i++ {
		for j := 0; j < len(nearbyTickets[i]); j++ {
			value := nearbyTickets[i][j]
			fit := 0
			for _, field := range fields {
				if field.withinRanges(value) {
					fit++
				}
			}

			if fit == 0 {
				sum += value
			}
		}
	}

	return sum
}
