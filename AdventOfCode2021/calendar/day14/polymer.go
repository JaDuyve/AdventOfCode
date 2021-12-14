package day14

import (
	"AdventOfCode2021/util/aocmath"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var expr = regexp.MustCompile("\\w+")

type Polymer struct {
	Template string
	rules    map[string]string
	freq     map[string]int
	maxSteps int
}

func New(problem string, maxSteps int) *Polymer {
	s := strings.Split(problem, "\n\n")

	ruleStrings := strings.Split(s[1], "\n")
	rules := make(map[string]string)

	for _, x := range ruleStrings {
		values := expr.FindAllString(x, -1)

		rules[values[0]] = values[1]
	}

	return &Polymer{
		s[0],
		rules,
		make(map[string]int),
		maxSteps,
	}
}

func (p *Polymer) ApplySteps() {
	begin := string(p.Template[0])
	fmt.Print(begin)
	for i := 0; i < len(p.Template)-2; i++ {
		p.applyStepsRec(0, p.Template[i:i+2])
		second := string(p.Template[i+2])
		fmt.Print(second)
	}
}

func (p *Polymer) applyStepsRec(depth int, pair string) {
	if depth == p.maxSteps {
		return
	}

	middle := p.rules[pair]
	p.freq[middle]++

	p.applyStepsRec(depth+1, string(pair[0])+middle)
	fmt.Printf(middle)
	p.applyStepsRec(depth+1, middle+string(pair[1]))
}

func (p Polymer) GetMinMaxDifference() int {
	min := math.MaxInt
	max := math.MinInt

	for _, v := range p.freq {
		min = aocmath.Min(min, v)
		max = aocmath.Max(max, v)
	}

	return max - min
}
