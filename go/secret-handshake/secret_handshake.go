/*
Package secret handshake

You and your fellow cohort of those in the "know" when it comes to
binary decide to come up with a secret "handshake".

```text
1 = wink
10 = double blink
100 = close your eyes
1000 = jump

10000 = Reverse the order of the operations in the secret handshake.
```

Given a decimal number, convert it to the appropriate sequence of events for a secret handshake.

Here's a couple of examples:

Given the input 3, the function would return the array
["wink", "double blink"] because 3 is 11 in binary.

Given the input 19, the function would return the array
["double blink", "wink"] because 19 is 10011 in binary.
Notice that the addition of 16 (10000 in binary)
has caused the array to be reversed.

Benchmarks:
BenchmarkHandshake-12             367724              3206 ns/op            1824 B/op         62 allocs/op
*/
package secret

// Handshake converts decimal number to the appropriate sequence of events for a secret handshake.
func Handshake(n uint) (res []string) {

	if n&1 > 0 {
		res = append(res, "wink")
	}
	if n&2 > 0 {
		res = append(res, "double blink")
	}
	if n&4 > 0 {
		res = append(res, "close your eyes")
	}
	if n&8 > 0 {
		res = append(res, "jump")
	}

	// reverse handshake
	if n&16 == 16 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}
