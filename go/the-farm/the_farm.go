package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalc FodderCalculator, cows int) (float64, error) {
	totalAmount, err := fodderCalc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	perCow := totalAmount * factor / float64(cows)

	return perCow, nil
}

func ValidateInputAndDivideFood(fodderCalc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalc, cows)
}

type InvalidCowsError struct {
	cows          int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.customMessage)
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows, "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows, "no cows don't need food"}
	default:
		return nil
	}
}
