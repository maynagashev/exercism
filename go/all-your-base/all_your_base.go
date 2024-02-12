package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// Check if all input digits are within the range of the input base
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	// Convert inputDigits to base 10
	base10 := 0
	for _, digit := range inputDigits {
		base10 = base10*inputBase + digit
	}
	//fmt.Printf("numbers: %+v, base10: %d, input base: %d, output base: %d\n", inputDigits, base10, inputBase, outputBase)

	// Handle the case where base10 is 0
	if base10 == 0 {
		return []int{0}, nil
	}

	// Convert base10 to outputBase
	outputDigits := []int{}
	for base10 > 0 {
		outputDigits = append([]int{base10 % outputBase}, outputDigits...)
		base10 /= outputBase
	}

	return outputDigits, nil
}
