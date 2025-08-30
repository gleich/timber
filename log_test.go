package timber

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
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

	level := Level{renderedMsg: "TEST"}
	args := []any{"hello", "world"}
	output := format(level, args...)

	parts := strings.SplitN(output.String(), " ", 3)
	if len(parts) != 3 {
		t.Fatalf("expected output to have 3 parts, got %d: %q", len(parts), output)
	}

	// Validate timestamp.
	if _, err := time.Parse(layout, parts[0]); err != nil {
		t.Errorf("timestamp %q is not valid: %v", parts[0], err)
	}

	// Validate level and message.
	if parts[1] != "TEST" {
		t.Errorf("expected log level to be %q, got %q", "TEST", parts[1])
	}
	if parts[2] != "hello world" {
		t.Errorf("expected message to be %q, got %q", "hello world", parts[2])
	}
}

func TestStackTraceLogging(t *testing.T) {
	// Capture err logger output and force stack printing.
	var buf bytes.Buffer

	globalLogger.mutex.Lock()
	oldShowStack := globalLogger.showStack
	oldErr := globalLogger.errOutput.logger
	globalLogger.showStack = true
	globalLogger.errOutput.logger = log.New(&buf, "", 0)
	globalLogger.mutex.Unlock()

	// Emit a non-fatal log so the process keeps running.
	ErrorMsg("unit-test")

	// Restore globals.
	globalLogger.mutex.Lock()
	globalLogger.showStack = oldShowStack
	globalLogger.errOutput.logger = oldErr
	globalLogger.mutex.Unlock()

	out := buf.String()
	parts := strings.SplitN(out, "\n", 2)
	if len(parts) < 2 {
		t.Fatalf("expected a newline then stack trace, got:\n%s", out)
	}
	trace := strings.TrimRight(parts[1], "\n")
	if trace == "" {
		t.Fatalf("expected non-empty stack trace, got empty")
	}

	lines := strings.Split(trace, "\n")

	// Each line should match: "#<n>. <function>()"
	re := regexp.MustCompile(`^#\d+\. .+\(\)$`)
	for _, ln := range lines {
		if !re.MatchString(ln) {
			t.Fatalf("stack frame does not match expected format: %q", ln)
		}
	}

	// Frames should be sequentially numbered starting at 1.
	for i, ln := range lines {
		want := "#" + strconv.Itoa(i+1) + ". "
		if !strings.HasPrefix(ln, want) {
			t.Fatalf("expected frame %d to start with %q, got %q", i+1, want, ln)
		}
	}

	// The last two runtime frames should have been trimmed.
	if strings.Contains(trace, "runtime.goexit") {
		t.Fatalf("expected runtime.goexit to be trimmed from the trace:\n%s", trace)
	}
	if strings.Contains(trace, "runtime.main") {
		t.Fatalf("expected runtime.main to be trimmed from the trace:\n%s", trace)
	}
}
