package timber

import (
	"fmt"
	"strings"
	"time"
)

func formatLog(level Level, v ...any) *strings.Builder {
	var out strings.Builder
	if globalLogger.displayTime {
		out.WriteString(time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat))
		out.WriteRune(' ')
	}
	out.WriteString(level.renderedMsg)
	out.WriteRune(' ')

	for i, item := range v {
		if i > 0 {
			out.WriteRune(' ')
		}
		fmt.Fprint(&out, item)
	}
	return &out
}

func logNormal(level Level, v ...any) {
	globalLogger.normalOutput.logger.Print(formatLog(level, v...).String())
}

func logFormatted(level Level, format string, v ...any) {
	logNormal(level, fmt.Sprintf(format, v...))
}

func logDurationNormal(level Level, start time.Time, v ...any) {
	logNormal(
		level,
		append(v, globalLogger.durationFormatter(time.Since(start)))...,
	)
}

func logError(err error, level Level, outputStack bool, formatted bool, v ...any) {
	var errorText string
	if err != nil {
		errorText = err.Error()
	}
	var out *strings.Builder
	if len(v) == 0 {
		out = formatLog(level, errorText)
	} else {
		out = formatLog(level, v...)
		if err != nil {
			out.WriteRune('\n')
			out.WriteString(err.Error())
		}
	}

	if outputStack {
		out.WriteRune('\n')
		callHeight := 4
		if formatted {
			callHeight++
		}
		stackTrace(out, callHeight)
	}
	globalLogger.errOutput.logger.Print(out.String())
}

func logDurationError(err error, level Level, outputStack bool, start time.Time, v ...any) {
	logError(
		err,
		level,
		outputStack,
		false,
		append(v, globalLogger.durationFormatter(time.Since(start)))...,
	)
}

func logErrorFormatted(err error, level Level, outputStack bool, format string, v ...any) {
	logError(err, level, outputStack, true, fmt.Sprintf(format, v...))
}
