package day14

import (
	"AdventOfCode2021/util/aocmath"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var expr = regexp.MustCompile("\\w+")

type Polymer struct {
	Template   string
	rules      map[string]string
	freq       map[string]int
	maxSteps   int
	history    map[string]map[string]int
	iterations int
}

func New(problem string, maxSteps int) *Polymer {
	s := strings.Split(problem, "\n\n")

	ruleStrings := strings.Split(s[1], "\n")
	rules := make(map[string]string)

	for _, x := range ruleStrings {
		values := expr.FindAllString(x, -1)

		rules[values[0]] = values[1]
	}

	template := s[0]
	freq := make(map[string]int)

	for i := range template {
		freq[template[i:i+1]]++
	}

	return &Polymer{
		template,
		rules,
		freq,
		maxSteps,
		make(map[string]map[string]int),
		0,
	}
}

func (p *Polymer) ApplySteps() {
	for i := 0; i < len(p.Template)-1; i++ {
		p.applyStepsRec(0, p.Template[i:i+2])
	}
}

func (p *Polymer) applyStepsRec(depth int, pair string) map[string]int {
	if depth == p.maxSteps {
		return make(map[string]int)
	}
	p.iterations++

	historyKey := strconv.Itoa(depth) + pair
	if value, ok := p.history[historyKey]; ok {
		for key, v := range value {
			p.freq[key] += v
		}
		return value
	}

	middle := p.rules[pair]
	p.freq[middle]++

	left := p.applyStepsRec(depth+1, string(pair[0])+middle)
	right := p.applyStepsRec(depth+1, middle+string(pair[1]))

	sum := make(map[string]int)

	for key, value := range left {
		sum[key] += value
	}

	for key, value := range right {
		sum[key] += value
	}

	sum[middle]++

	p.history[historyKey] = sum

	return sum
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
