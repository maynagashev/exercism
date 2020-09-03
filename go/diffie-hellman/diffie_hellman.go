/*
Package diffiehellman contains functions for generating private -> public -> secret keys

Benchmarks:

1. Using *big.Int.Exp()
BenchmarkPrivateKey
BenchmarkPrivateKey-12           1928329               619 ns/op             496 B/op          6 allocs/op
BenchmarkPublicKey
BenchmarkPublicKey-12               4096            293639 ns/op            6179 B/op         26 allocs/op
BenchmarkNewPair
BenchmarkNewPair-12                 1878            582472 ns/op            6676 B/op         32 allocs/op
BenchmarkSecretKey
BenchmarkSecretKey-12               4092            290920 ns/op            6011 B/op         23 allocs/op

Tags: random generator
*/
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey returns a random bigint in range (2, p)
func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	// generating [0, p-2)
	r, err := rand.Int(rand.Reader, big.NewInt(0).Sub(p, two))
	if err != nil {
		panic(err)
	}
	// shifting generated value to [2, p)
	return r.Add(r, two)
}

// PublicKey returns g^a (mod p)
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return modexp(big.NewInt(g), private, p)
}

// NewPair simply return a private and public key for given primes
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey calculates the shared secret from public and private keys
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return modexp(public2, private1, p)
}

func modexp(base, exp, modulus *big.Int) *big.Int {
	return new(big.Int).Exp(base, exp, modulus)
}
