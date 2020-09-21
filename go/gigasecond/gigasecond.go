package gigasecond

import (
	"math"
	"time"
)

// Adds  10^9 seconds to a given time.
func AddGigasecond(t time.Time) time.Time {
	ts := t.Unix() + int64(math.Pow(10, 9))
	return time.Unix(ts, 0)
}
