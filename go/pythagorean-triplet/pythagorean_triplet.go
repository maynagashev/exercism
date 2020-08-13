/*
Package pythagorean contains various generators of Pythagorean triplets

A Pythagorean triplet is a set of three natural numbers, {a, b, c}, for
which,
```text
a**2 + b**2 = c**2
```
and such that,
```text
a < b < c
```
*/
package pythagorean

// Triplet defines a Pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (res []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if (a*a + b*b) == c*c {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) (res []Triplet) {
	for _, t := range Range(1, p-2) {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return res
}
