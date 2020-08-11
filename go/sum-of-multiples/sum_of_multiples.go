/*
Package summultiples contains function for calculating sum of multiples.

Given a number, find the sum of all the unique multiples of particular numbers up to
but not including that number.

If we list all the natural numbers below 20 that are multiples of 3 or 5,
we get 3, 5, 6, 9, 10, 12, 15, and 18.

The sum of these multiples is 78.
*/
package summultiples

// SumMultiples calculates the sum of all the natural below limit that are multiples of given divisors
func SumMultiples(limit int, divisors ...int) (sum int) {
	for mp := 1; mp < limit; mp++ {
		for _, d := range divisors {
			if d != 0 && mp%d == 0 {
				sum += mp
				break // for uniqueness of mp
			}
		}
	}
	return
}
