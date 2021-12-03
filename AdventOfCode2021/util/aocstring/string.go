package aocstring

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func SplitToIntArray(str string, delimiter string) []int {
	stringArray := strings.Split(str, delimiter)

	intArray := make([]int, len(stringArray))

	for i, numberString := range stringArray {
		number, err := strconv.Atoi(numberString)

		if err != nil {
			log.Fatal(err)
		}

		intArray[i] = number
	}

	return intArray
}

func ReadProblemString(filePath string) string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
