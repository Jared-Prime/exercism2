// Package gigasecond extends the standard libary for "time" with additional
// constants and functions for very large time intervals, ex. 10^9 seconds aka "gigasecond"
// http://exercism.io/submissions/7dd2cbf4a0c444ed8ffb62441fe48470
package gigasecond

import "time"

const (
	// Gigasecond is 10^9 (1,000,000,000) seconds
	Gigasecond = time.Second * 1000000000
)

// AddGigasecond adds 10^9 seconds to the provided time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
