// Package calculate the Hamming Distance between two DNA strands.
// Calculate the Hamming Distance between two DNA strands.
//  Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing
//  into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!
package hamming

import "errors"

// Distance returns the hamming distance of two strings. The comparison is done
// at the character level, not the underlying bits.
func Distance(a, b string) (int, error) {
	// Set the initial distance to 0
	distance := 0

	// The Hamming distance is only defined for sequences of equal length, so
	// an attempt to calculate it between sequences of different lengths should not work.
	if len(a) != len(b) {
		return -1, errors.New("input must have the same length")
	}

	// If the characters are different the spacing increases.
	for i := range a {
		if b[i] != a[i] {
			distance++
		}
	}

	return distance, nil
}
