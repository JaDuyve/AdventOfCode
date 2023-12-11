package main

import (
	"log"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2023, 11, session)
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
	res2 := part2(input, 1_000_000)
	us = time.Since(startTime).Microseconds()
	log.Printf("Result: %d in %dμs\n", res2, us)
}

// part one
func part1(input string) int {
	s := Space{
		Data:     ExpandSpace(input),
		Galaxies: make([]utils.Tuple[int, int], 0),
	}

	s.findAllGalaxies()

	return s.CalculateTotalDistance()
}

func ExpandSpace(input string) [][]uint8 {
	galaxyInput := strings.Split(input, "\n")

	width := len(galaxyInput[0])
	for i := len(galaxyInput) - 1; i >= 0; i-- {
		if false == strings.ContainsFunc(galaxyInput[i], func(x rune) bool {
			return x != '.'
		}) {
			galaxyInput = append(galaxyInput[:i+1], galaxyInput[i:]...)
			galaxyInput[i] = strings.Repeat(".", width)
		}
	}

	d := make([][]uint8, len(galaxyInput))
	for r := 0; r < len(galaxyInput); r++ {
		d[r] = []byte(galaxyInput[r])
	}

	for x := width - 1; x >= 0; x-- {
		found := false
		for y := 0; y < len(d); y++ {
			if d[y][x] != '.' {
				found = true
			}
		}

		if found == false {
			for yy := 0; yy < len(d); yy++ {
				d[yy] = append(d[yy][:x+1], d[yy][x:]...)
				d[yy][x] = '.'
			}
		}
	}

	return d
}

// part two
func part2(input string, spaceMulti int) int {
	lines := strings.Split(input, "\n")
	data := make([][]uint8, len(lines))
	for i := 0; i < len(lines); i++ {
		data[i] = []byte(lines[i])
	}

	s := Space{
		Data:     data,
		Galaxies: make([]utils.Tuple[int, int], 0),
	}
	s.findAllGalaxies()
	s.Expand(spaceMulti)

	return s.CalculateTotalDistance()
}

type Space struct {
	Data     [][]uint8
	Galaxies []utils.Tuple[int, int]
}

func (s *Space) String() string {
	builder := strings.Builder{}

	for i := 0; i < len(s.Data); i++ {
		builder.Write(s.Data[i])
		builder.WriteByte('\n')
	}

	return builder.String()
}

func (s *Space) findAllGalaxies() {
	for y := 0; y < len(s.Data); y++ {
		for x := 0; x < len(s.Data[y]); x++ {
			if s.Data[y][x] == '#' {
				s.Galaxies = append(s.Galaxies, utils.Tuple[int, int]{A: x, B: y})
			}
		}
	}
}

func (s *Space) CalculateTotalDistance() int {
	sum := 0

	for i := 0; i < len(s.Galaxies); i++ {
		p1 := s.Galaxies[i]
		for j := i + 1; j < len(s.Galaxies); j++ {
			p2 := s.Galaxies[j]
			distance := utils.ManhattenDistance(p1, p2)
			sum += distance
		}
	}

	return sum
}

func (s *Space) Expand(distMult int) {
	newGalaxies := make([]utils.Tuple[int, int], len(s.Galaxies))
	for j, g := range s.Galaxies {
		newGalaxies[j] = utils.Tuple[int, int]{
			A: g.A,
			B: g.B,
		}
	}

	for y := 0; y < len(s.Data); y++ {
		empty := true
		for x := 0; x < len(s.Data[y]); x++ {
			if s.Data[y][x] == '#' {
				empty = false
			}
		}

		if empty {
			for i := 0; i < len(s.Galaxies); i++ {
				if s.Galaxies[i].B > y {
					newGalaxies[i].B += distMult - 1
				}
			}
		}
	}

	for x := 0; x < len(s.Data[0]); x++ {
		empty := true
		for y := 0; y < len(s.Data); y++ {
			if s.Data[y][x] == '#' {
				empty = false
			}
		}

		if empty {
			for i := 0; i < len(s.Galaxies); i++ {
				if s.Galaxies[i].A > x {
					newGalaxies[i].A += distMult - 1
				}
			}
		}
	}

	s.Galaxies = newGalaxies
}
