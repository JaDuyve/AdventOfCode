package day07

import (
	"AdventOfCode2020/utils/conv"
	"AdventOfCode2020/utils/files"
	"go/types"
	"regexp"
	"strings"
)

var parentRegex = regexp.MustCompile("\\w+ \\w+")
var childRegex = regexp.MustCompile("(\\d) (\\w+ \\w+)")

type Bag struct {
	color    string
	contains []ChildBag
}

type ChildBag struct {
	color  string
	amount int
}

func Part1() int {
	input := files.ReadFile("calendar/day07/input")
	return solvePart1(input)

}

func solvePart1(puzzleInput string) int {
	rules := strings.Split(puzzleInput, "\n")
	bags := buildBags(rules)

	color := "shiny gold"

	visitedBags := make(map[string]types.Nil)
	countOuterBagsOf(color, &visitedBags, &bags)
	return len(visitedBags)
}

func buildBags(rules []string) map[string]Bag {
	bags := make(map[string]Bag)

	for _, rule := range rules {
		bag := *createBag(rule)
		bags[bag.color] = bag
	}

	return bags
}

func createBag(rule string) *Bag {
	return &Bag{
		color:    getParentColor(rule),
		contains: getChildBags(rule),
	}
}

func getParentColor(rule string) string {
	return parentRegex.FindString(rule)
}

func getChildBags(rule string) []ChildBag {
	childBags := make([]ChildBag, 0)

	matches := childRegex.FindAllStringSubmatch(rule, -1)

	for _, match := range matches {
		childBags = append(childBags, ChildBag{
			color:  match[2],
			amount: conv.ToInt(match[1]),
		})
	}
	return childBags
}

func countOuterBagsOf(color string, visited *map[string]types.Nil, bags *map[string]Bag) {
	foundBags := getOuterBagsContaining(color, bags)

	for bag := range *foundBags {
		(*visited)[bag] = types.Nil{}
		countOuterBagsOf(bag, visited, bags)
	}
}

func getOuterBagsContaining(color string, bags *map[string]Bag) *map[string]types.Nil {
	outerBags := make(map[string]types.Nil)

	for _, outer := range *bags {
		for _, inner := range outer.contains {
			if inner.color == color {
				outerBags[outer.color] = types.Nil{}
			}
		}
	}

	return &outerBags
}
