package timber

import (
	"fmt"
	"strings"
	"time"
)

func formatLog(level Level, msg string, start time.Time, attributes []Value) string {
	if globalLogger.structured.enabled {
		return formatStructured(level, msg, start, attributes)
	}
	return formatPlain(level, msg, start, attributes)
}

func formatStructured(level Level, msg string, start time.Time, attributes []Value) string {
	if !start.IsZero() {
		attributes = append([]Value{{"duration", formatDuration(time.Since(start))}}, attributes...)
	}
	out := make([]string, 0, 3+len(attributes))
	out = append(out,
		time.Now().UTC().Format(globalLogger.structured.timeFormat),
		fmt.Sprintf("level=%q", level.Message),
		fmt.Sprintf("msg=%q", msg),
	)
	if len(attributes) > 0 {
		fmtAttributes := make([]string, 0, len(attributes))
		for _, attribute := range attributes {
			fmtAttributes = append(
				fmtAttributes,
				fmt.Sprintf(`%s="%v"`, attribute.Key, attribute.Data),
			)
		}
		out = append(out, strings.Join(fmtAttributes, " "))
	}
	return strings.Join(out, " ")
}

func formatPlain(level Level, msg string, start time.Time, attributes []Value) string {
	if !start.IsZero() {
		msg = fmt.Sprintf("%s (%s)", msg, formatDuration(time.Since(start)))
	}
	out := make([]string, 0, 3)
	out = append(out,
		time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat),
		level.renderedMsg,
		msg,
	)
	if len(attributes) > 0 {
		fmtAttributes := make([]string, 0, len(attributes))
		for _, attribute := range attributes {
			fmtAttributes = append(
				fmtAttributes,
				fmt.Sprintf("%s: %v", attribute.Key, attribute.Data),
			)
		}
		out = append(out, "["+strings.Join(fmtAttributes, ", ")+"]")
	}
	return strings.Join(out, " ")
}

func outputNormal(s string) {
	globalLogger.normalOutput.logger.Print(s)
}

func logNormal(level Level, msg string, attributes []Value) {
	outputNormal(formatLog(level, msg, time.Time{}, attributes))
}

func logDurationNormal(level Level, start time.Time, msg string, attributes []Value) {
	outputNormal(formatLog(level, msg, start, attributes))
}

func outputError(
	level Level,
	err error,
	msg string,
	start time.Time,
	attributes []Value,
	outputStack bool,
) {
	structured := globalLogger.structured.enabled
	var errText string
	if err != nil {
		errText = err.Error()
		if structured {
			attributes = append([]Value{{"error", errText}}, attributes...)
		}
	}
	out := formatLog(level, msg, start, attributes)
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
	attributes []Value,
	outputStack bool,
) {
	outputError(level, err, msg, time.Time{}, attributes, outputStack)
}

func logDurationError(
	level Level,
	err error,
	start time.Time,
	msg string,
	attributes []Value,
	outputStack bool,
) {
	outputError(level, err, msg, start, attributes, outputStack)
}
