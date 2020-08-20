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

Append to new slice
BenchmarkHandshake-12             367724              3206 ns/op            1824 B/op         62 allocs/op

Remove from filled slice
BenchmarkHandshake-12             703927              1709 ns/op            2048 B/op         32 allocs/op


*/
package secret

// Handshake converts decimal number to the appropriate sequence of events for a secret handshake.
func Handshake(n uint) []string {
	var res = []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
	}
	var shift uint8 = 0

	if n&1 == 0 {
		res = remove(res, 0)
		shift++
	}
	if n&2 == 0 {
		res = remove(res, 1-shift)
		shift++
	}
	if n&4 == 0 {
		res = remove(res, 2-shift)
		shift++
	}
	if n&8 == 0 {
		res = remove(res, 3-shift)
	}

	// reverse handshake
	if n&16 == 16 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}

// remove removes item from slice by index
func remove(slice []string, i uint8) []string {
	return append(slice[:i], slice[i+1:]...)
}
