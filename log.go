package timber

import (
	"fmt"
	"os"
	"runtime"
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
		out.WriteString(fmt.Sprint(item))
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

	out := formatLog(level, v...)
	if err != nil {
		out.WriteRune('\n')
		out.WriteString(err.Error())
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

func stackTrace(builder *strings.Builder) {
	pcs := make([]uintptr, 64)
	n := runtime.Callers(3, pcs) // skip runtime.Callers, stackTrace, logError/Error*
	if n == 0 {
		return
	}
	pcs = pcs[:n]

	if len(pcs) < 2 {
		return
	}
	pcs = pcs[:len(pcs)-2]

	fr := runtime.CallersFrames(pcs)
	var frames []runtime.Frame
	for {
		f, more := fr.Next()
		frames = append(frames, f)
		if !more {
			break
		}
	}

	for i := 0; i < len(frames); i++ {
		if i+1 < len(frames) {
			fmt.Fprintf(
				builder,
				"#%d. %s() [%s:%d]\n",
				i+1,
				frames[i].Function,
				frames[i+1].File,
				frames[i+1].Line,
			)
		} else {
			fmt.Fprintf(builder, "#%d. %s()\n", i+1, frames[i].Function)
		}
	}
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
