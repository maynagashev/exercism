// Package gigasecond contains method for work with gigasecond duration.
// A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond calculates the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow(10, 9)) * time.Second)
}
