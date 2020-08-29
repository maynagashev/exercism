/*
Package perfect contains classification functions for natural numbers.
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
	case alq == n:
		return ClassificationPerfect, nil
	case alq < n:
		return ClassificationDeficient, nil
	case alq > n:
		return ClassificationAbundant, nil
	}

	return Unclassified, nil
}

func aliquotSum(n int64) (sum int64) {
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}
