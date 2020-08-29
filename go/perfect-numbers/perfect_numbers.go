/*
Package perfect contains classification functions for natural numbers.

Benchmarks:
1. Not optimized aliquotSum
BenchmarkClassify
BenchmarkClassify-12                   2         713268213 ns/op               4 B/op          0 allocs/op

2. Optimized aliquotSum
BenchmarkClassify
BenchmarkClassify-12                9475            126192 ns/op               0 B/op          0 allocs/op
*/
package perfect

import "errors"

// Classification return type
type Classification uint8

// Class constants
const (
	Unclassified Classification = iota
	ClassificationDeficient
	ClassificationAbundant
	ClassificationPerfect
)

// ErrOnlyPositive error when we expecting only natural number
var ErrOnlyPositive error = errors.New("expected only natural number")

// Classify determines if a number is perfect, abundant, or deficient based on
// Nicomachus' (60 - 120 CE) classification scheme for natural numbers
func Classify(n int64) (class Classification, err error) {

	if n < 1 {
		return Unclassified, ErrOnlyPositive
	}

	// calculate
	alq := aliquotSum(n)

	// classify
	switch {
	case alq < n || n == 1:
		return ClassificationDeficient, nil
	case alq == n:
		return ClassificationPerfect, nil
	case alq > n:
		return ClassificationAbundant, nil
	}

	return Unclassified, nil
}

func aliquotSum(n int64) (sum int64) {
	sum = 1
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if n/i != i {
				sum += n / i
			}
		}
	}
	return
}
