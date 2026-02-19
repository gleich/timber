package timber

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   time.Duration
		want string
	}{
		{"zero", 0, "0s"},

		{"1ms", 1 * time.Millisecond, "[1ms]"},
		{"999ms", 999 * time.Millisecond, "[999ms]"},
		{"negative-ms", -250 * time.Millisecond, "[-250ms]"},

		{"1us", 1 * time.Microsecond, "[1µs]"},
		{"999us", 999 * time.Microsecond, "[999µs]"},
		{"negative-us", -12 * time.Microsecond, "[-12µs]"},

		{"1s", 1 * time.Second, "[1.0s]"},
		{"1.5s", 1500 * time.Millisecond, "[1.5s]"},
		{"negative-seconds", -1500 * time.Millisecond, "[-1.5s]"},

		{"2min", 2*time.Minute + 0*time.Second, "[2min 0.0s]"},
		{"2min 5.4s", 2*time.Minute + 5400*time.Millisecond, "[2min 5.4s]"},
		{"1h", 1*time.Hour + 0*time.Second, "[1h 0.0s]"},
		{"1h 2min 3.1s", 1*time.Hour + 2*time.Minute + 3100*time.Millisecond, "[1h 2min 3.1s]"},
		{"1d", 24*time.Hour + 0*time.Second, "[1d 0.0s]"},
		{"1d 3h 0.9s", 24*time.Hour + 3*time.Hour + 900*time.Millisecond, "[1d 3h 0.9s]"},
		{
			"2d 3h 4min 5.6s",
			48*time.Hour + 3*time.Hour + 4*time.Minute + 5600*time.Millisecond,
			"[2d 3h 4min 5.6s]",
		},
		{
			"negative-combo",
			-(48*time.Hour + 3*time.Hour + 4*time.Minute + 5600*time.Millisecond),
			"[-2d 3h 4min 5.6s]",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := formatDuration(tt.in); got != tt.want {
				t.Fatalf("formatDuration(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFormatDuration_Properties(t *testing.T) {
	t.Parallel()

	if got := formatDuration(-1 * time.Nanosecond); got[:2] != "[-" {
		t.Fatalf("expected negative duration to start with \"[-\", got %q", got)
	}

	cases := []time.Duration{
		1 * time.Nanosecond,
		1 * time.Microsecond,
		1 * time.Millisecond,
		1 * time.Second,
		2*time.Minute + 3*time.Second,
	}
	for _, d := range cases {
		got := formatDuration(d)
		if got[0] != '[' || got[len(got)-1] != ']' {
			t.Fatalf("expected %v to be bracketed, got %q", d, got)
		}
	}
}
