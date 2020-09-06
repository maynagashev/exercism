/*
Package clock contains type Clock for handling times without dates.

Benchmarks

1. With two attributes in a struct AND immutable Clock.Add, Clock.Subtract:
BenchmarkAddMinutes
BenchmarkAddMinutes-12          21001855                49.0 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     26232820                49.4 ns/op             0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        30703069                38.7 ns/op             0 B/op          0 allocs/op

2. With two attributes in a struct (hours and minutes)
BenchmarkAddMinutes
BenchmarkAddMinutes-12          40211384                27.5 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     38940823                29.0 ns/op             0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        31202306                38.6 ns/op             0 B/op          0 allocs/op

3. Only minutes in a struct (not using getters Clock.Hours, Clock.Minutes in "stringer" method) - BEST RESULT
BenchmarkAddMinutes
BenchmarkAddMinutes-12          113015074                9.74 ns/op            0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     123854264                9.59 ns/op            0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        219707516                5.45 ns/op            0 B/op          0 allocs/op

4. Only minutes in a struct AND using getters Clock.Hours, Clock.Minutes in Clock.String method - MORE READABLE
BenchmarkAddMinutes
BenchmarkAddMinutes-12          100000000               10.1 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     100000000               10.1 ns/op             0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        100000000               10.5 ns/op             0 B/op          0 allocs/op

5. Removed getters
BenchmarkAddMinutes
BenchmarkAddMinutes-12          119078599               10.3 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     149693394                8.95 ns/op            0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        227830412                5.31 ns/op            0 B/op          0 allocs/op

6. Removed normalizeMinutes function and named constant
BenchmarkAddMinutes
BenchmarkAddMinutes-12          96593484                10.4 ns/op             0 B/op          0 allocs/op
BenchmarkSubtractMinutes
BenchmarkSubtractMinutes-12     141532027                9.21 ns/op            0 B/op          0 allocs/op
BenchmarkCreateClocks
BenchmarkCreateClocks-12        224030209                5.35 ns/op            0 B/op          0 allocs/op
*/
package clock

import "fmt"

// Clock type
type Clock struct {
	minutes int
}

// New creates new Clock structure from arbitrary integer values of hours and minutes.
// Can handle negative values and roll overs.
func New(h int, m int) Clock {
	minutes := (h*60 + m) % (24 * 60)
	if minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{minutes}
}

// String converts Clock to a formatted string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add minutes to given Clock instance
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

// Subtract minutes from given Clock instance
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}
