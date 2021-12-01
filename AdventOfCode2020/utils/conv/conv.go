package conv

import "strconv"

func ToIntSlice(slice []string) []int {
	sliceToReturn := make([]int, 0, len(slice))

	for _, a := range slice {
		convertedString, err := strconv.Atoi(a)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToInt(str string) int {
	convertedString, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return convertedString
}
