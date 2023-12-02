package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsValidBirthYear_UpperLimit_Valid(t *testing.T) {
	assert.True(t, isValidBirthYear("2002"), "Year 2002 should be valid!")
}

func Test_IsValidBirthYear_UpperLimit_Invalid(t *testing.T) {
	assert.False(t, isValidBirthYear("2003"), "Year 2003 should be invalid!")
}

func Test_IsValidBirthYear_LowerLimit_Valid(t *testing.T) {
	assert.True(t, isValidBirthYear("1920"), "Year 1920 should be valid!")
}

func Test_IsValidBirthYear_LowerLimit_Invalid(t *testing.T) {
	assert.False(t, isValidBirthYear("1919"), "Year 1919 should be invalid!")
}

func Test_IsValidBirthYear_RandomYear_Valid(t *testing.T) {
	assert.True(t, isValidBirthYear("1972"), "Year 1972 should be valid!")
}

func Test_IsValidIssueYear_UpperLimit_Valid(t *testing.T) {
	assert.True(t, isValidIssueYear("2020"), "Year 2020 should be valid!")
}

func Test_IsValidIssueYear_UpperLimit_Invalid(t *testing.T) {
	assert.False(t, isValidIssueYear("2021"), "Year 2021 should be invalid!")
}

func Test_IsValidIssueYear_LowerLimit_Valid(t *testing.T) {
	assert.True(t, isValidIssueYear("2010"), "Year 2010 should be valid!")
}

func Test_IsValidIssueYear_LowerLimit_Invalid(t *testing.T) {
	assert.False(t, isValidIssueYear("2009"), "Year 2009 should be invalid!")
}

func Test_IsValidIssueYear_RandomYear_Valid(t *testing.T) {
	assert.True(t, isValidIssueYear("2015"), "Year 2015 should be valid!")
}

func Test_IsValidExpirationYear_UpperLimit_Valid(t *testing.T) {
	assert.True(t, isValidExpirationYear("2030"), "Year 2030 should be valid!")
}

func Test_IsValidExpirationYear_UpperLimit_Invalid(t *testing.T) {
	assert.False(t, isValidExpirationYear("2031"), "Year 2031 should be invalid!")
}

func Test_IsValidExpirationYear_LowerLimit_Valid(t *testing.T) {
	assert.True(t, isValidExpirationYear("2020"), "Year 2020 should be valid!")
}

func Test_IsValidExpirationYear_LowerLimit_Invalid(t *testing.T) {
	assert.False(t, isValidExpirationYear("2019"), "Year 2019 should be invalid!")
}

func Test_IsValidExpirationYear_RandomYear_Valid(t *testing.T) {
	assert.True(t, isValidExpirationYear("2025"), "Year 2025 should be valid!")
}

func Test_IsValidHeight_WithCM_Valid(t *testing.T) {
	assert.True(t, isValidHeight("160cm"), "Height \"190cm\" should be valid")
}

func Test_IsValidHeight_WithIN_Valid(t *testing.T) {
	assert.True(t, isValidHeight("60in"), "Height \"60in\" should be valid")
}

func Test_IsValidHeight_WithoutMetric_Invalid(t *testing.T) {
	assert.False(t, isValidHeight("60"), "Height value \"60\" has no metric -> should be invalid!")
}

func Test_IsValidEyeColor_ExistingColors_Valid(t *testing.T) {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range colors {
		assert.Truef(t, isValidEyeColor(color), "Eye color \"%s\" should be valid", color)
	}
}

func Test_IsValidEyeColor_NonExistingColor_Invalid(t *testing.T) {
	assert.False(t, isValidEyeColor("gmt"), "Eye color \"gmt\" should be invalid")
}

func Test_IsValidHairColor_Valid(t *testing.T) {
	assert.True(t, isValidHairColor("#123abc"), "Hair color \"#123abc\" should be valid")
}

func Test_IsValidHairColor_OutCharRange_Invalid(t *testing.T) {
	assert.False(t, isValidHairColor("#123abz"), "Hair color \"#123abz\" should be invalid")
}

func Test_IsValidHairColor_WithoutHashtag_Invalid(t *testing.T) {
	assert.False(t, isValidHairColor("123abc"), "Hair color \"123abc\" should be invalid")
}

func Test_IsValidPassportId_Valid(t *testing.T) {
	assert.True(t, isValidPassportId("123456789"), "Passport Id \"123456789\" should be valid")
}

func Test_IsValidPassportId_WithLeadingZeros_Valid(t *testing.T) {
	assert.True(t, isValidPassportId("000000001"), "Passport Id \"000000001\" should be valid")
}

func Test_IsValidPassportId_TooShort_Invalid(t *testing.T) {
	assert.False(t, isValidPassportId("12345678"), "Passport Id \"12345678\" should be invalid")
}

func Test_IsValidPassportId_TooLong_Invalid(t *testing.T) {
	assert.False(t, isValidPassportId("0123456789"), "Passport Id \"0123456789\" should be invalid")
}
