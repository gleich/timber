package timber

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func formatLog(level Level, v ...any) *strings.Builder {
	var out strings.Builder
	out.WriteString(time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat))
	out.WriteRune(' ')
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
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	globalLogger.normalOutput.logger.Print(formatLog(level, v...).String())
}

func logFormatted(level Level, format string, v ...any) {
	logNormal(level, fmt.Sprintf(format, v...))
}

func logError(err error, level Level, outputStack bool, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()

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
		stackTrace(out)
	}
	globalLogger.errOutput.logger.Print(out.String())
}

func logErrorFormatted(err error, level Level, outputStack bool, format string, v ...any) {
	logError(err, level, outputStack, fmt.Sprintf(format, v...))
}

// Output a DEBUG-level message
func Debug(v ...any) {
	logNormal(globalLogger.levels.Debug, v...)
}

// Output a formatted DEBUG-level message
func Debugf(format string, v ...any) {
	logFormatted(globalLogger.levels.Debug, format, v...)
}

// Output a DONE-level message
func Done(v ...any) {
	logNormal(globalLogger.levels.Done, v...)
}

// Output a formatted DONE-level message
func Donef(format string, v ...any) {
	logFormatted(globalLogger.levels.Done, format, v...)
}

// Output a INFO-level message
func Info(v ...any) {
	logNormal(globalLogger.levels.Info, v...)
}

// Output a formatted INFO-level message
func Infof(format string, v ...any) {
	logFormatted(globalLogger.levels.Info, format, v...)
}

// Output a WARN-level message
func Warning(v ...any) {
	logNormal(globalLogger.levels.Warning, v...)
}

// Output a formatted WARN-level message
func Warningf(format string, v ...any) {
	logFormatted(globalLogger.levels.Warning, format, v...)
}

// Output an ERROR log message with information about the error
func Error(err error, v ...any) {
	logError(err, globalLogger.levels.Error, globalLogger.showErrorStack, v...)
}

// Output a formatted ERROR-level message with information about the error
func Errorf(err error, format string, v ...any) {
	logErrorFormatted(err, globalLogger.levels.Error, globalLogger.showErrorStack, format, v...)
}

// Output a ERROR-level message
func ErrorMsg(v ...any) {
	logError(nil, globalLogger.levels.Error, globalLogger.showErrorStack, v...)
}

// Output a ERROR-level message
func ErrorMsgf(format string, v ...any) {
	logErrorFormatted(nil, globalLogger.levels.Error, globalLogger.showErrorStack, format, v...)
}

// Output a FATAL-level message with information about the error
func Fatal(err error, v ...any) {
	logError(err, globalLogger.levels.Fatal, globalLogger.showFatalStack, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a formatted FATAL-level message with information about the error
func Fatalf(err error, format string, v ...any) {
	logErrorFormatted(err, globalLogger.levels.Fatal, globalLogger.showFatalStack, format, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a FATAL-level message
func FatalMsg(v ...any) {
	logError(nil, globalLogger.levels.Fatal, globalLogger.showFatalStack, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a formatted FATAL-level message
func FatalMsgf(format string, v ...any) {
	logErrorFormatted(nil, globalLogger.levels.Fatal, globalLogger.showFatalStack, format, v...)
	os.Exit(globalLogger.fatalExitCode)
}
