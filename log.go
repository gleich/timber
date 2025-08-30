package timber

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func format(level Level, v ...any) *strings.Builder {
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
	globalLogger.normalOutput.logger.Print(format(level, v...).String())
}

func logError(err error, level Level, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()

	out := format(level, v...)
	if err != nil {
		out.WriteRune('\n')
		out.WriteString(err.Error())
	}
	if globalLogger.showStack {
		out.WriteRune('\n')
		stackTrace(out)
	}
	globalLogger.errOutput.logger.Print(out.String())
}

func stackTrace(builder *strings.Builder) {
	programCounters := make([]uintptr, 1024)
	n := runtime.Callers(3, programCounters)
	if n == 0 {
		return
	}
	programCounters = programCounters[:n]

	// exclude the last two frames (runtime.main, runtime.goexit)
	if len(programCounters) >= 2 {
		programCounters = programCounters[:len(programCounters)-2]
	} else {
		return
	}

	frames := runtime.CallersFrames(programCounters)
	for i := 1; ; i++ {
		frame, more := frames.Next()
		fmt.Fprintf(builder, "#%d. %s()\n", i, frame.Function)
		if !more {
			break
		}
	}
}

// Output a DEBUG log message
func Debug(v ...any) {
	logNormal(globalLogger.levels.Debug, v...)
}

// Output a DONE log message
func Done(v ...any) {
	logNormal(globalLogger.levels.Done, v...)
}

// Output an INFO log message
func Info(v ...any) {
	logNormal(globalLogger.levels.Info, v...)
}

// Output a WARN log message
func Warning(v ...any) {
	logNormal(globalLogger.levels.Warning, v...)
}

// Output an ERROR log message with information about the error
func Error(err error, v ...any) {
	logError(err, globalLogger.levels.Error, v...)
}

// Output an ERROR log message
func ErrorMsg(v ...any) {
	logError(nil, globalLogger.levels.Error, v...)
}

// Output a FATAL log message with information about the error
func Fatal(err error, v ...any) {
	logError(err, globalLogger.levels.Fatal, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a FATAL log message
func FatalMsg(v ...any) {
	logError(nil, globalLogger.levels.Fatal, v...)
	os.Exit(globalLogger.fatalExitCode)
}
