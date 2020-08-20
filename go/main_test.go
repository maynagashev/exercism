package main

import "testing"

func BenchmarkHandshake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			Handshake(j)
		}
	}
}
