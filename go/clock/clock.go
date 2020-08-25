/*
Package clock contains type Clock for handling times without dates.

Benchmarks

1. Immutable add, subtract:
BenchmarkAddMinutes
BenchmarkAddMinutes-12          21001855                49.0 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     26232820                49.4 ns/op             0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        30703069                38.7 ns/op             0 B/op          0 allocs/op

2. Current
BenchmarkAddMinutes
BenchmarkAddMinutes-12          40211384                27.5 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     38940823                29.0 ns/op             0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        31202306                38.6 ns/op             0 B/op          0 allocs/op
*/
package clock

import "fmt"

// Clock type
type Clock struct {
	h int
	m int
}

// New creates new Clock structure from arbitrary integer values.
// Can handle negative values and roll overs.
func New(h int, m int) Clock {
	h, m = transformInput(h, m)
	return Clock{h, m}
}

// String converts Clock to a formatted string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add minutes to given Clock instance
func (c Clock) Add(m int) Clock {
	c.h, c.m = transformInput(c.h, c.m+m)
	return c
	//return New(c.h, c.m+m) // immutable
}

// Subtract minutes from given Clock instance
func (c Clock) Subtract(m int) Clock {
	c.h, c.m = transformInput(c.h, c.m-m)
	return c
	//return New(c.h, c.m-m) // immutable
}

// transformInput transforms arbitrary input to proper values, handles roll overs and negative values
func transformInput(h int, m int) (int, int) {

	// negative minutes
	if m < 0 {
		h = h - 1 + m/60
		m = m % 60
		m += 60
	}

	// minutes rolls over
	if m > 59 {
		h += m / 60
		m = m % 60
	}

	// negative hours
	if h < 0 {
		h = h % 24
		h += 24
	}

	// hours rolls over
	if h > 23 {
		h = h % 24
	}

	return h, m
}
