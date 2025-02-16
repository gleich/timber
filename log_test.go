package timber

import (
	"strings"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// Use a fixed time format (ISO 8601) without spaces to simplify parsing.
	const layout = "2006-01-02T15:04:05Z07:00"

	// Update the global logger's time configuration.
	globalLogger.mutex.Lock()
	globalLogger.timeFormat = layout
	globalLogger.timezone = time.UTC
	globalLogger.mutex.Unlock()

	level := Level{
		renderedMsg: "TEST",
	}
	args := []any{"hello", "world"}
	output := format(level, args...)

	// Expected format: "<timestamp> <renderedMsg> <arguments joined by a space>"
	// Split the output into three parts:
	// 1. The timestamp string.
	// 2. The log level rendered message.
	// 3. The remaining message (joined arguments).
	parts := strings.SplitN(output.String(), " ", 3)
	if len(parts) != 3 {
		t.Fatalf("expected output to have 3 parts, got %d: %q", len(parts), output)
	}

	// Validate the timestamp by parsing it using the same layout.
	timestampStr := parts[0]
	if _, err := time.Parse(layout, timestampStr); err != nil {
		t.Errorf("timestamp %q is not valid: %v", timestampStr, err)
	}

	// Validate that the log level rendered message is as expected.
	if parts[1] != "TEST" {
		t.Errorf("expected log level to be %q, got %q", "TEST", parts[1])
	}

	// Validate the message part (arguments).
	expectedMsg := "hello world"
	if parts[2] != expectedMsg {
		t.Errorf("expected message to be %q, got %q", expectedMsg, parts[2])
	}
}
