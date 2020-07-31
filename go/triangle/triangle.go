/*
Package triangle contains functions for work with triangles.
*/
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64 {a, b, c}
	for _, side := range sides {
		if side == 0 {
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
