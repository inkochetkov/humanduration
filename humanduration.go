// Package humanduration formats time.Duration into human-readable strings.
//
// Unlike time.Duration.String(), which produces machine-friendly output
// like "24h0m0s", humanduration focuses on readability for logs, CLI
// output, and user-facing messages.
//
// The package has zero dependencies outside the standard library and
// exposes only three functions.
package humanduration

import (
	"fmt"
	"strings"
	"time"
)

const (
	week = 7 * 24 * time.Hour
	day  = 24 * time.Hour
)

type unit struct {
	d     time.Duration
	long  string
	short string
}

var units = []unit{
	{week, "week", "w"},
	{day, "day", "d"},
	{time.Hour, "hour", "h"},
	{time.Minute, "minute", "m"},
	{time.Second, "second", "s"},
	{time.Millisecond, "millisecond", "ms"},
}

// Format returns a human-readable representation of d using at most two
// units, for example "1 hour 30 minutes", "2 days", or "500 milliseconds".
// Zero returns "0 seconds". Negative durations are prefixed with "-".
func Format(d time.Duration) string { return FormatN(d, 2) }

// FormatN is like Format but keeps at most n units. If n < 1 it is
// treated as 1.
func FormatN(d time.Duration, n int) string {
	if n < 1 {
		n = 1
	}
	if d == 0 {
		return "0 seconds"
	}
	neg := d < 0
	if neg {
		d = -d
	}
	var parts []string
	for _, u := range units {
		if len(parts) >= n {
			break
		}
		if d < u.d {
			continue
		}
		count := d / u.d
		d -= count * u.d
		parts = append(parts, fmt.Sprintf("%d %s%s", count, u.long, plural(count)))
	}
	if len(parts) == 0 {
		// d is smaller than 1ms.
		return "0 seconds"
	}
	out := strings.Join(parts, " ")
	if neg {
		out = "-" + out
	}
	return out
}

// FormatShort returns a compact human-readable form using at most two
// units and short suffixes, for example "1h 30m", "2d", or "500ms".
func FormatShort(d time.Duration) string {
	if d == 0 {
		return "0s"
	}
	neg := d < 0
	if neg {
		d = -d
	}
	var parts []string
	for _, u := range units {
		if len(parts) >= 2 {
			break
		}
		if d < u.d {
			continue
		}
		count := d / u.d
		d -= count * u.d
		parts = append(parts, fmt.Sprintf("%d%s", count, u.short))
	}
	if len(parts) == 0 {
		return "0s"
	}
	out := strings.Join(parts, " ")
	if neg {
		out = "-" + out
	}
	return out
}

func plural(n time.Duration) string {
	if n == 1 {
		return ""
	}
	return "s"
}