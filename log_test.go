package timber

import (
	"regexp"
	"testing"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// TestFormat verifies that the format function produces an output string that
// begins with a timestamp in the expected format, followed by the rendered log level
// and then the concatenated log message.
func TestFormat(t *testing.T) {
	logger.timezone = time.UTC
	logger.timeFormat = "2006-01-02 15:04:05"

	style := lipgloss.NewStyle()

	level := logLevel("INFO")
	msg1 := "hello"
	msg2 := "world"

	output := format(level, style, msg1, msg2)

	if len(output) < 20 {
		t.Fatalf("output too short: %q", output)
	}
	timestampPart := output[:19]
	if _, err := time.Parse("2006-01-02 15:04:05", timestampPart); err != nil {
		t.Errorf("timestamp %q is not in expected format: %v", timestampPart, err)
	}

	expectedSuffix := "INFO hello world"
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} INFO hello world$`)
	if !re.MatchString(output) {
		t.Errorf("output %q does not match expected format", output)
	}

	actualSuffix := output[20:] // skip the 19-char timestamp and the following space
	if actualSuffix != expectedSuffix {
		t.Errorf("expected message %q, got %q", expectedSuffix, actualSuffix)
	}
}
