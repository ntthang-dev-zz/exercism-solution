// Package darts returns the earned points in a single toss of a Darts game
// In our particular instance of the game, the target rewards with 4 different amounts of points, depending on where the dart lands:
//  If the dart lands outside the target, player earns no points (0 points).
//  If the dart lands in the outer circle of the target, player earns 1 point.
//  If the dart lands in the middle circle of the target, player earns 5 points.
//  If the dart lands in the inner circle of the target, player earns 10 points.
package darts

// Function that given a point in the target (defined by its `real` cartesian
// coordinates `x` and `y`), returns the correct amount earned by a dart landing in that point.
func Score(x, y float64) int {
	distance := (x * x) + (y * y)

	switch {
	// Inner circle : 1 radius
	case distance <= 1:
		return 10
	// Middle circle : 5 radius
	case distance <= 25:
		return 5
	// Outer circle : 10 radius
	case distance <= 100:
		return 1
	}
	// Outside the target
	return 0
}
