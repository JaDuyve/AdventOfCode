package day10

import (
	"AdventOfCode2020/utils/conv"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_CalculateJoltDifference_InputExample1(t *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	adapters := conv.ToIntSlice(strings.Split(input, "\n"))
	adapters = makeAdapterChain(adapters)

	expected := []int{0, 7, 0, 5}

	assert.Equal(t, expected, calculateJoltDifference(adapters))
}

func Test_CalculateJoltDifference_InputExample2(t *testing.T) {
	input := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"
	adapters := conv.ToIntSlice(strings.Split(input, "\n"))
	adapters = makeAdapterChain(adapters)

	expected := []int{0, 22, 0, 10}

	assert.Equal(t, expected, calculateJoltDifference(adapters))
}

func Test_SolvePart1_TestInputExample1(t *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"

	expected := 7 * 5
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart1_TestInputExample2(t *testing.T) {
	input := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"

	expected := 22 * 10
	assert.Equal(t, expected, solvePart1(input))
}

func Test_SolvePart2_TestInputExample1(t *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"

	expected := 8
	assert.Equal(t, expected, solvePart2(input))
}

func Test_SolvePart2_TestInputExample2(t *testing.T) {
	input := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"

	expected := 19208
	assert.Equal(t, expected, solvePart2(input))
}
