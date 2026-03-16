package timber

import (
	"fmt"
	"strings"
	"time"
)

func formatLog(level Level, msg string, start time.Time, vals []Value) string {
	if globalLogger.structured.enabled {
		return formatStructured(level, msg, start, vals)
	}
	return formatPlain(level, msg, start, vals)
}

func formatStructured(level Level, msg string, start time.Time, vals []Value) string {
	if !start.IsZero() {
		vals = append([]Value{{"duration", formatDuration(time.Since(start))}}, vals...)
	}
	out := make([]string, 0, 3+len(vals))
	out = append(out,
		time.Now().UTC().Format(globalLogger.structured.timeFormat),
		fmt.Sprintf("level=%q", level.Message),
		fmt.Sprintf("msg=%q", msg),
	)
	if len(vals) > 0 {
		fmtValues := make([]string, 0, len(vals))
		for _, attribute := range vals {
			fmtValues = append(
				fmtValues,
				fmt.Sprintf(`%s="%v"`, attribute.Key, attribute.Data),
			)
		}
		out = append(out, strings.Join(fmtValues, " "))
	}
	return strings.Join(out, " ")
}

func formatPlain(level Level, msg string, start time.Time, vals []Value) string {
	if !start.IsZero() {
		msg = fmt.Sprintf("%s (%s)", msg, formatDuration(time.Since(start)))
	}
	out := make([]string, 0, 3)
	out = append(out,
		time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat),
		level.renderedMsg,
		msg,
	)
	if len(vals) > 0 {
		fmtValues := make([]string, 0, len(vals))
		for _, attribute := range vals {
			fmtValues = append(
				fmtValues,
				fmt.Sprintf("%s: %v", attribute.Key, attribute.Data),
			)
		}
		out = append(out, "["+strings.Join(fmtValues, ", ")+"]")
	}
	return strings.Join(out, " ")
}

func outputNormal(s string) {
	globalLogger.normalOutput.logger.Print(s)
}

func logNormal(level Level, msg string, vals []Value) {
	outputNormal(formatLog(level, msg, time.Time{}, vals))
}

func logDurationNormal(level Level, start time.Time, msg string, vals []Value) {
	outputNormal(formatLog(level, msg, start, vals))
}

func outputError(
	level Level,
	err error,
	msg string,
	start time.Time,
	vals []Value,
	outputStack bool,
) {
	structured := globalLogger.structured.enabled
	var errText string
	if err != nil {
		errText = err.Error()
		if structured {
			vals = append([]Value{{"error", errText}}, vals...)
		}
	}
	out := formatLog(level, msg, start, vals)
	if err != nil && !structured {
		out += "\n" + errText
	}
	if outputStack {
		stackTrace(&out, 5)
	}
	globalLogger.errOutput.logger.Print(out)
}

func logError(
	level Level,
	err error,
	msg string,
	vals []Value,
	outputStack bool,
) {
	outputError(level, err, msg, time.Time{}, vals, outputStack)
}

func logDurationError(
	level Level,
	err error,
	start time.Time,
	msg string,
	vals []Value,
	outputStack bool,
) {
	outputError(level, err, msg, start, vals, outputStack)
}
