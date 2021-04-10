// Package acronym convert a phrase to its acronym.
//  Techies love their TLA (Three Letter Acronyms)!
//  Help generate some jargon by writing a program that converts a long name
//  like Portable Network Graphics to its acronym (PNG).
package acronym

import "strings"

// Abbreviate to create acronym of a string.
func Abbreviate(s string) string {
	// Create output variables
	var out string

	// Replace the characters "-" and "_"
	// in the string with a space
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	words := strings.Fields(s)

	// Perform to get the first characters of the words in the
	// string into the output string
	for i := range words {
		out += string(words[i][0])
	}

	// Performs output string capitalization
	return strings.ToUpper(out)
}
