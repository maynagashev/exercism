/*
Package complexnumbers implements the following operations:
- addition, subtraction, multiplication and division of two complex numbers,
- conjugate, absolute value, exponent of a given complex number.
*/
package complexnumbers

import "math"

// Number complex number
type Number struct {
	a float64
	b float64
}

// Real returns the real part of a complex number
func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.a + n2.a, n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.a - n2.a, n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.a*n2.a - n1.b*n2.b, n1.a*n2.b + n1.b*n2.a}
}

func (n Number) Times(factor float64) Number {
	return Number{n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.a*n2.a + n2.b*n2.b
	return Number{(n1.a*n2.a + n1.b*n2.b) / denominator, (n1.b*n2.a - n1.a*n2.b) / denominator}
}

func (n Number) Conjugate() Number {
	return Number{n.a, -n.b}
}

// Abs The absolute value of a complex number z = a + b * i is a real number |z| = sqrt(a^2 + b^2)
func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

// Exp returns the exponent of a given complex number
// Raising e to a complex exponent can be expressed as e^(a + i * b) = e^a * e^(i * b), the last term of which is given by Euler's formula e^(i * b) = cos(b) + i * sin(b).
func (n Number) Exp() Number {
	return Number{math.Exp(n.a) * math.Cos(n.b), math.Exp(n.a) * math.Sin(n.b)}
}
