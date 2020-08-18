/*
Package grains about grains count on chessboard.

There once was a wise servant who saved the life of a prince. The king
promised to pay whatever the servant could dream up. Knowing that the
king loved chess, the servant told the king he would like to have grains
of wheat. One grain on the first square of a chess board, with the number
of grains doubling on each successive square.

There are 64 squares on a chessboard (where square 1 has one grain, square 2 has two grains, and so on).

1. math.Pow with fmt.Errorf()
BenchmarkSquare
BenchmarkSquare-12       1990354               580 ns/op             160 B/op          8 allocs/op
BenchmarkTotal
BenchmarkTotal-12       37896012                31.1 ns/op             0 B/op          0 allocs/op

2. "for" loop with fmt.Errorf()
BenchmarkSquare
BenchmarkSquare-12       2420720               489 ns/op             160 B/op          8 allocs/op
BenchmarkTotal
BenchmarkTotal-12       34725582                34.4 ns/op             0 B/op          0 allocs/op

3. bitwise multiplication with fmt.Errorf()
BenchmarkSquare
BenchmarkSquare-12       2563512               445 ns/op             160 B/op          8 allocs/op
BenchmarkTotal
BenchmarkTotal-12       1000000000               0.252 ns/op           0 B/op          0 allocs/op

4. bitwise multiplication without fmt.Errorf()
BenchmarkSquare
BenchmarkSquare-12      38860022                28.4 ns/op             0 B/op          0 allocs/op
BenchmarkTotal
BenchmarkTotal-12       1000000000               0.503 ns/op           0 B/op          0 allocs/op

*/
package grains

import (
	"errors"
)

// Square calculates how many grains were on a given square.
func Square(n int) (res uint64, err error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid square number")
	}
	return 1 << (n - 1), nil
}

// Total calculates the total number of grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
