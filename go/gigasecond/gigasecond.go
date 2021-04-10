// Package gigasecond Given a moment, determine the moment that would be after a gigasecond has passed.
// A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
