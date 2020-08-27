package main

func main() {
	Sieve(20)
}
func Sieve(limit int) (primes []int) {

	for i := 2; i < limit; i++ {
		primes = append(primes, i)
	}
	return
}
