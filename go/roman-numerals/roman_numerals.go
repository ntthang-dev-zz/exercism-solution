// Package romannumerals to convert from normal numbers to Roman Numerals.
// 	The Romans were a clever bunch. They conquered most of Europe and ruled
// 	it for hundreds of years. They invented concrete and straight roads and
// 	even bikinis. One thing they never discovered though was the number
// 	zero. This made writing and dating extensive histories of their exploits
// 	slightly more challenging, but the system of numbers they came up with
// 	is still in use today. For example the BBC uses Roman numerals to date
// 	their programmes.
//
//	The Romans wrote numbers using letters - I, V, X, L, C, D, M. (notice
//	these letters have lots of straight lines and are hence easy to hack
//	into stone tablets).
package romannumerals

import "errors"

// ToRomanNumeral convert from normal numbers to Roman Numerals.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("number out of range")
	}
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for arabic >= conversion.value {
			roman += conversion.digit
			arabic -= conversion.value
		}
	}
	return roman, nil
}
