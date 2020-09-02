package diffiehellman

import "math/big"

// PrivateKey returns a random bigint in range (2, p)
func PrivateKey(p *big.Int) *big.Int {
	return p
}

// PublicKey returns g^a (mod p)
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return p
}

// NewPair simply return a private and public key for given primes
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	return
}

// SecretKey calculates the shared secret from public and private keys
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return private1
}