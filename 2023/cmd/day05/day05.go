package main

import (
	"cmp"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 05, session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Part One ---")
	startTime := time.Now().Local()
	res1 := part1(input)
	us := time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res1, us)

	log.Println("--- Part Two ---")
	startTime = time.Now().Local()
	res2 := part2(input)
	us = time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res2, us)
}

// part one
func part1(input string) int {
	mappingGroupStrings := utils.SplitToStringGroups(input, "\n\n", "\n")
	seeds := utils.SplitTo[int](strings.Split(mappingGroupStrings[0][0], ": ")[1], " ", utils.ParseInt)
	mappingGroups := make([][][]int, len(mappingGroupStrings)-1)

	for i, mg := range mappingGroupStrings[1:] {
		mappings := make([][]int, len(mg)-1)
		for j, m := range mg[1:] {
			mappings[j] = utils.SplitTo[int](m, " ", utils.ParseInt)
		}

		mappingGroups[i] = mappings
	}

	min := math.MaxInt

	for _, seed := range seeds {
		currentValue := seed
		for i := 0; i < len(mappingGroups); i++ {
			for j := 0; j < len(mappingGroups[i]); j++ {
				m := mappingGroups[i][j]
				if m[1] <= currentValue && currentValue < m[1]+m[2] {
					currentValue = currentValue - m[1] + m[0]
					break
				}
			}
		}

		if currentValue < min {
			min = currentValue
		}
	}

	return min
}

func comp(a, b Range) int {
	return cmp.Compare(a.Start, b.Start)
}

// part two
func part2(input string) int {
	mappingGroupStrings := utils.SplitToStringGroups(input, "\n\n", "\n")
	seeds := utils.SplitTo[int](strings.Split(mappingGroupStrings[0][0], ": ")[1], " ", utils.ParseInt)
	rangeSeeds := make([]Range, len(seeds)/2)
	for s := 0; s < len(seeds); s += 2 {
		rangeSeeds[s/2] = Range{
			Start: seeds[s],
			End:   seeds[s] + seeds[s+1] - 1,
		}
	}
	mappingGroups := make([][][]int, len(mappingGroupStrings)-1)

	for i, mg := range mappingGroupStrings[1:] {
		mappings := make([][]int, len(mg)-1)
		for j, m := range mg[1:] {
			mappings[j] = utils.SplitTo[int](m, " ", utils.ParseInt)
		}

		mappingGroups[i] = mappings
	}

	min := math.MaxInt

	for _, rs := range rangeSeeds {
		originalRanges := []Range{
			rs,
		}
		for i := 0; i < len(mappingGroups); i++ {
			mappedRanges := []Range{}

			for _, or := range originalRanges {

				mappedGroupRanges := []Range{}
				stateRange := []Range{}
				for j := 0; j < len(mappingGroups[i]); j++ {
					m := mappingGroups[i][j]
					start := m[1]
					end := m[1] + m[2]

					if (or.Start >= start && or.End <= end) || (or.Start <= start && or.End < end) || (or.Start > start && or.End <= end) {
						newStart := utils.Max(start, or.Start)
						newEnd := utils.Min(end, or.End)
						newRange := Range{
							Start: newStart - start + m[0],
							End:   (newEnd - start) + m[0] - 1,
						}

						stateRange = append(stateRange, Range{Start: newStart, End: newEnd})
						mappedGroupRanges = append(mappedGroupRanges, newRange)
					}
				}

				sort.Slice(stateRange, func(i, j int) bool {
					return stateRange[i].Start < stateRange[j].Start
				})

				if len(stateRange) == 0 {
					mappedRanges = append(mappedRanges, or)
				} else {
					start := or.Start

					for x := 0; x < len(stateRange); x++ {
						if start < stateRange[x].Start {
							mappedRanges = append(mappedRanges, Range{Start: start, End: stateRange[x].Start - 1})
							start = stateRange[x].End
						}
					}

					if or.End > stateRange[len(stateRange)-1].End {
						mappedRanges = append(mappedRanges, Range{Start: stateRange[len(stateRange)-1].End + 1, End: or.End})
					}
				}

				mappedRanges = append(mappedRanges, mappedGroupRanges...)
			}

			originalRanges = mappedRanges
		}

		for _, r := range originalRanges {
			if r.Start < min {
				min = r.Start
			}
		}
	}

	return min
}

type Range struct {
	Start int
	End   int
}
