package slices

import (
	"fmt"
)

func StringFilter(slice []string, filter func(string) bool) []string {
	resSlice := make([]string, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}

func IntFilter(slice []int, filter func(int) bool) []int {
	resSlice := make([]int, 0)

	for _, element := range slice {
		if filter(element) {
			resSlice = append(resSlice, element)
		}
	}

	return resSlice
}

func ToTwoDimArray(slice []string) [][]rune {
	retSlice := make([][]rune, 0, len(slice))

	for _, str := range slice {
		retSlice = append(retSlice, []rune(str))
	}

	return retSlice
}

func Print2DimRuneArray(slice [][]rune) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%c\n", str)
	}

	print(result)
}

func Print2DimIntArray(slice [][]int) {
	var result string

	for _, str := range slice {
		result += fmt.Sprintf("%v\n", str)
	}

	print(result)
}

func Remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func CopyString4Dim(src [][][][]string) [][][][]string {
	dst := make([][][][]string, 0, len(src))

	for i := 0; i < len(src); i++ {
		dst = append(dst, CopyString3Dim(src[i]))
	}

	return dst
}

func CopyString3Dim(src [][][]string) [][][]string {
	dst := make([][][]string, 0, len(src))

	for i := 0; i < len(src); i++ {
		dst = append(dst, CopyString2Dim(src[i]))
	}

	return dst
}

func CopyString2Dim(src [][]string) [][]string {
	dst := make([][]string, 0, len(src))

	for i := 0; i < len(src); i++ {
		a := make([]string, len(src[i]))
		copy(a, src[i])
		dst = append(dst, a)
	}

	return dst
}
