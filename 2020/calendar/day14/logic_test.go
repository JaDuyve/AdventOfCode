package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetIndexValuePair(t *testing.T) {
	a := assert.New(t)
	operationString := "mem[65461] = 458355100"
	expectedIndex := 65461
	expectedValue := 458355100

	extractedIndex, extractedValue := getIndexValuePair(operationString)
	a.Equal(expectedIndex, extractedIndex)
	a.Equal(expectedValue, extractedValue)
}

func Test_GenerateIndexes_Mask1(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	index := 42

	expectedIndexes := []int{26, 27, 58, 59}
	assert.Equal(t, expectedIndexes, generateIndexes(mask, index))
}

func Test_GenerateIndexes_Mask2(t *testing.T) {
	mask := "00000000000000000000000000000000X0XX"
	index := 26

	expectedIndexes := []int{16, 17, 18, 19, 24, 25, 26, 27}
	assert.Equal(t, expectedIndexes, generateIndexes(mask, index))
}
