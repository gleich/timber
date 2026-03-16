package timber

import (
	"fmt"
	"strings"
	"time"
)

func insertDuration(msg *string, attributes *[]Value, start time.Time) {
	dur := formatDuration(time.Since(start))
	if globalLogger.structured.enabled {
		*attributes = append([]Value{{"duration", dur}}, *attributes...)
	} else {
		*msg = fmt.Sprintf("%s (%s)", *msg, dur)
	}
}

func formatDuration(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	sign := ""
	if d < 0 {
		sign = "-"
		d = -d
	}

	if d < time.Second {
		if d >= time.Millisecond {
			ms := d.Milliseconds()
			if ms == 0 && d > 0 {
				ms = 1
			}
			return fmt.Sprintf("%s%dms", sign, ms)
		}

		us := d.Microseconds()
		if us == 0 && d > 0 {
			us = 1
		}
		return fmt.Sprintf("%s%dµs", sign, us)
	}

	var parts []string

	if days := d / (24 * time.Hour); days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
		d %= 24 * time.Hour
	}

	if hours := d / time.Hour; hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
		d %= time.Hour
	}

	if minutes := d / time.Minute; minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dmin", minutes))
		d %= time.Minute
	}

	sec := float64(d) / float64(time.Second)
	parts = append(parts, fmt.Sprintf("%.1fs", sec))

	return fmt.Sprintf("%s%s", sign, strings.Join(parts, " "))
}
