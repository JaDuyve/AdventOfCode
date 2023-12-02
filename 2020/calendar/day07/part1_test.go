package day07

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_GetParent(t *testing.T) {
	ruleString := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	expected := "light red"
	assert.Equalf(t, expected, getParentColor(ruleString), "Wrong parent")
}

func Test_GetChildren(t *testing.T) {
	ruleString := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	expected := []ChildBag{
		{
			color:  "bright white",
			amount: 1,
		},
		{
			color:  "muted yellow",
			amount: 2,
		},
	}

	result := getChildBags(ruleString)
	assert.Truef(t, reflect.DeepEqual(expected, result), "Result getChildBags should be %v but was %v for %s.", expected, result, ruleString)
}
