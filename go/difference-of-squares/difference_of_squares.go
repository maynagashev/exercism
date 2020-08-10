/*
Package diffsquares contains functions for calculating difference between the square of the sum and the sum of the squares
*/
package diffsquares

// SquareOfSum calculates (1 + 2 + ... + n)²
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates 1² + 2² + ... + n²
func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return
}

// Difference calculates difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
