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
	// Capture err logger output.
	var buf bytes.Buffer

	// Save and override global settings safely.
	globalLogger.mutex.Lock()
	oldShowStack := globalLogger.showStack
	oldErrLogger := globalLogger.errOutput.logger
	globalLogger.showStack = true
	globalLogger.errOutput.logger = log.New(&buf, "", 0)
	globalLogger.mutex.Unlock()

	// Emit an error (non-fatal) so we don't exit the test process.
	ErrorMsg("unit-test")

	// Restore globals.
	globalLogger.mutex.Lock()
	globalLogger.showStack = oldShowStack
	globalLogger.errOutput.logger = oldErrLogger
	globalLogger.mutex.Unlock()

	out := buf.String()
	if !strings.Contains(out, "\n#1. ") {
		t.Fatalf("expected stack trace lines in log output, got:\n%s", out)
	}

	// Extract the trace portion: everything after the first newline.
	nl := strings.IndexByte(out, '\n')
	if nl == -1 || nl+1 >= len(out) {
		t.Fatalf("log output missing newline/trace: %q", out)
	}
	trace := out[nl+1:]
	lines := strings.Split(strings.TrimRight(trace, "\n"), "\n")
	if len(lines) == 0 {
		t.Fatalf("expected at least one stack frame, got none")
	}

	// 1) Each frame should match: "#<n>. <file>:<line> -> <func>()"
	re := regexp.MustCompile(`^#\d+\. .+:\d+ -> .+\(\)$`)
	for _, ln := range lines {
		if !re.MatchString(ln) {
			t.Fatalf("stack frame doesn't match expected format: %q", ln)
		}
	}

	// 2) Frames should be sequentially numbered starting at 1.
	for i, ln := range lines {
		wantPrefix := "#" + strconv.Itoa(i+1) + ". "
		if !strings.HasPrefix(ln, wantPrefix) {
			t.Fatalf("expected frame %d to start with %q, got %q", i+1, wantPrefix, ln)
		}
	}

	// 3) The last two runtime frames should have been trimmed.
	if strings.Contains(trace, "runtime.goexit") {
		t.Fatalf("expected runtime.goexit to be trimmed from the trace:\n%s", trace)
	}
}
