package aocstring

import (
	"strings"
)

func ExtractFieldOfString(pairs []string, delimiter string) map[string]string {
	var extractedFields = make(map[string]string)

	for _, pair := range pairs {
		splitPair := strings.Split(pair, delimiter)

		key := splitPair[0]
		value := splitPair[1]

		if _, ok := extractedFields[key]; ok == false {
			extractedFields[key] = value
		}
	}

	return extractedFields
}
