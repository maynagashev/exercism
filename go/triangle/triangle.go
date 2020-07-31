/*
Package triangle contains functions for work with triangles.
*/
package triangle

import "math"

// Kind is resulting type of triangles
type Kind int

/*
NaT - not a triangle
Equ - equilateral
Iso - isosceles
Sca - scalene
*/
const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
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
