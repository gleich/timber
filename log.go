package timber

import (
	"fmt"
	"os"
	"runtime/debug"
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
	globalLogger.logger.Print(format(level, v...).String())
}

func logError(err error, level Level, v ...any) {
	out := format(level, v...)
	if err != nil && globalLogger.showStack {
		out.WriteRune('\n')
		out.WriteString(err.Error())
		out.WriteRune('\n')
		out.WriteString(string(debug.Stack()))
	} else if err != nil {
		out.WriteRune('\n')
		out.WriteString(err.Error())
	}
	globalLogger.errLogger.Print(out.String())
}

// Output a DEBUG log message
func Debug(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Debug, v...)
}

// Output a DONE log message
func Done(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Done, v...)
}

// Output a INFO log message
func Info(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Info, v...)
}

// Output a WARN log message
func Warning(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Warning, v...)
}

// Output a ERROR log message with information about the error
func Error(err error, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(err, globalLogger.levels.Error, v...)
}

// Output a ERROR log message
func ErrorMsg(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(nil, globalLogger.levels.Error, v...)
}

// Output a FATAL log message with information about the error
func Fatal(err error, v ...any) {
	globalLogger.mutex.RLock()
	logError(err, globalLogger.levels.Fatal, v...)
	globalLogger.mutex.RUnlock()
	os.Exit(globalLogger.fatalExitCode)
}

// Output a FATAL log message
func FatalMsg(v ...any) {
	globalLogger.mutex.RLock()
	logError(nil, globalLogger.levels.Fatal, v...)
	globalLogger.mutex.RUnlock()
	os.Exit(globalLogger.fatalExitCode)
}
