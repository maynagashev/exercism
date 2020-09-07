/*
Package clock contains type Clock for handling times without dates.
*/
package clock

import "fmt"

// Clock type
type Clock struct {
	minutes int
}

// New creates new Clock structure from arbitrary integer values of hours and minutes.
// Can handle negative values and roll overs.
func New(hours int, minutes int) Clock {
	m := hours*60 + minutes
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
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
