package timber

import (
	"fmt"
	"strings"
	"time"
)

func formatLog(level Level, msg string, start time.Time, attrs []Attr) string {
	if globalLogger.structured.enabled {
		return formatStructured(level, msg, start, attrs)
	}
	return formatPlain(level, msg, start, attrs)
}

func formatStructured(level Level, msg string, start time.Time, attrs []Attr) string {
	if !start.IsZero() {
		attrs = append([]Attr{{"duration", formatDuration(time.Since(start))}}, attrs...)
	}
	out := make([]string, 0, 3+len(attrs))
	out = append(out,
		time.Now().UTC().Format(globalLogger.structured.timeFormat),
		fmt.Sprintf("level=%q", level.Message),
		fmt.Sprintf("msg=%q", msg),
	)
	if len(attrs) > 0 {
		fmtValues := make([]string, 0, len(attrs))
		for _, attribute := range attrs {
			fmtValues = append(
				fmtValues,
				fmt.Sprintf(`%s="%v"`, attribute.Key, attribute.Value),
			)
		}
		out = append(out, strings.Join(fmtValues, " "))
	}
	return strings.Join(out, " ")
}

func formatPlain(level Level, msg string, start time.Time, attrs []Attr) string {
	if !start.IsZero() {
		msg = fmt.Sprintf("%s (%s)", msg, formatDuration(time.Since(start)))
	}
	out := make([]string, 0, 3)
	out = append(out,
		time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat),
		level.renderedMsg,
		msg,
	)
	if len(attrs) > 0 {
		fmtValues := make([]string, 0, len(attrs))
		for _, attribute := range attrs {
			fmtValues = append(
				fmtValues,
				fmt.Sprintf("%s: %v", attribute.Key, attribute.Value),
			)
		}
		out = append(out, "["+strings.Join(fmtValues, ", ")+"]")
	}
	return strings.Join(out, " ")
}

func outputNormal(s string) {
	globalLogger.normalOutput.logger.Print(s)
}

func logNormal(level Level, msg string, attrs []Attr) {
	outputNormal(formatLog(level, msg, time.Time{}, attrs))
}

func logDurationNormal(level Level, start time.Time, msg string, attrs []Attr) {
	outputNormal(formatLog(level, msg, start, attrs))
}

func outputError(
	level Level,
	err error,
	msg string,
	start time.Time,
	vals []Attr,
	outputStack bool,
) {
	structured := globalLogger.structured.enabled
	var errText string
	if err != nil {
		errText = err.Error()
		if structured {
			vals = append([]Attr{{"error", errText}}, vals...)
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
	attrs []Attr,
	outputStack bool,
) {
	outputError(level, err, msg, time.Time{}, attrs, outputStack)
}

func logDurationError(
	level Level,
	err error,
	start time.Time,
	msg string,
	attrs []Attr,
	outputStack bool,
) {
	outputError(level, err, msg, start, attrs, outputStack)
}
