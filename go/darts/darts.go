package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y) // Pythagorean theorem to calc distance between (0, 0) and (x, y)
	switch {
	case distance > 10:
		return 0 // Outside the target
	case distance > 5:
		return 1 // Outer circle
	case distance > 1:
		return 5 // Middle circle
	default:
		return 10 // Inner circle
	}
}
