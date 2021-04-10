// Package strand given a DNA strand, return its RNA complement (per RNA transcription).
package strand

import "strings"

// ToRNA given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement
func ToRNA(dna string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(dna)
}
