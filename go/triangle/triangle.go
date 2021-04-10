// Package triangle determine if a triangle is equilateral, isosceles, or scalene.
// 	An equilateral triangle has all three sides the same length.
// 	An isosceles triangle has at least two sides the same length. (It is sometimes specified as having exactly
// 	two sides the same length, but for the purposes of this exercise we'll say at least two.)
// 	A scalene triangle has all sides of different lengths.
package triangle

import "math"

// Kind is the triangle kind. KindFromSides() returns this type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             //s equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides tells if a triangle is equilateral, isosceles or scalene.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
