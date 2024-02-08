package prime

func Factors(n int64) []int64 {
	var factors []int64
	for i := int64(2); n > 1; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}
