// Package gigasecond contains method and constant for work with gigasecond duration.
// A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

import (
	"time"
)

// Gigasecond constant with package specific duration
const Gigasecond = 1e9 * time.Second

// AddGigasecond calculates the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
