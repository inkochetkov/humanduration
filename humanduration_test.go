package humanduration

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		d    time.Duration
		want string
	}{
		{0, "0 seconds"},
		{time.Second, "1 second"},
		{2 * time.Second, "2 seconds"},
		{time.Minute + 30*time.Second, "1 minute 30 seconds"},
		{time.Hour + 30*time.Minute, "1 hour 30 minutes"},
		{25 * time.Hour, "1 day 1 hour"},
		{7 * 24 * time.Hour, "1 week"},
		{9 * 24 * time.Hour, "1 week 2 days"},
		{500 * time.Millisecond, "500 milliseconds"},
		{-90 * time.Second, "-1 minute 30 seconds"},
		{100 * time.Microsecond, "0 seconds"},
		{time.Nanosecond, "0 seconds"},
	}
	for _, c := range cases {
		if got := Format(c.d); got != c.want {
			t.Errorf("Format(%v) = %q, want %q", c.d, got, c.want)
		}
	}
}

func TestFormatN(t *testing.T) {
	cases := []struct {
		d    time.Duration
		n    int
		want string
	}{
		{90 * time.Second, 1, "1 minute"},
		{90 * time.Second, 2, "1 minute 30 seconds"},
		{90 * time.Second, 3, "1 minute 30 seconds"},
		{25 * time.Hour, 1, "1 day"},
		{25 * time.Hour, 2, "1 day 1 hour"},
		{90 * time.Second, 0, "1 minute"},
		{90 * time.Second, -5, "1 minute"},
	}
	for _, c := range cases {
		if got := FormatN(c.d, c.n); got != c.want {
			t.Errorf("FormatN(%v, %d) = %q, want %q", c.d, c.n, got, c.want)
		}
	}
}

func TestFormatShort(t *testing.T) {
	cases := []struct {
		d    time.Duration
		want string
	}{
		{0, "0s"},
		{time.Second, "1s"},
		{90 * time.Second, "1m 30s"},
		{time.Hour + 30*time.Minute, "1h 30m"},
		{25 * time.Hour, "1d 1h"},
		{7 * 24 * time.Hour, "1w"},
		{9 * 24 * time.Hour, "1w 2d"},
		{500 * time.Millisecond, "500ms"},
		{-90 * time.Second, "-1m 30s"},
	}
	for _, c := range cases {
		if got := FormatShort(c.d); got != c.want {
			t.Errorf("FormatShort(%v) = %q, want %q", c.d, got, c.want)
		}
	}
}