package armstrong

import (
	"fmt"
	"math"
)

// IsNumber An Armstrong number is a number that is the sum of its own digits
// each raised to the power of the number of digits.
func IsNumber(n int) bool {
	power := len(fmt.Sprint(n))
	sum := 0
	temp := n
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(power)))
		temp /= 10
	}
	return sum == n
}
