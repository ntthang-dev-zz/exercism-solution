// Package raindrops is to convert a number into a string that contains raindrop sounds
// corresponding to certain potential factors. A factor is a number that evenly divides
// into another number, leaving no remainder.
package raindrops

import (
	"bytes"
	"strconv"
)

// Case rules of raindrops are that if a given number
var translations = []struct {
	Num  int
	Word string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert takes an integer value and "translates" it to a particular raindrop sound by
// determining the integer's prime factorization.
func Convert(amount int) string {
	var b bytes.Buffer

	for _, trans := range translations {
		if amount%trans.Num == 0 {
			b.WriteString(trans.Word)
		}
	}

	if b.String() == "" {
		b.WriteString(strconv.Itoa(amount))
	}

	return b.String()
}
