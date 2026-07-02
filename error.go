package timber

import (
	"os"
	"time"
)

// Output an ERROR-level message with information about the error
func Error(err error, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(err, globalLogger.levels.Error, globalLogger.showErrorStack, false, v...)
}

// Output an ERROR-level message since a certain time with information about the error
func ErrorSince(err error, start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(err, globalLogger.levels.Error, globalLogger.showErrorStack, start, v...)
}

// Output a formatted ERROR-level message with information about the error
func Errorf(err error, format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logErrorFormatted(err, globalLogger.levels.Error, globalLogger.showErrorStack, format, v...)
}

// Output a ERROR-level message
func ErrorMsg(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(nil, globalLogger.levels.Error, globalLogger.showErrorStack, false, v...)
}

// Output an ERROR-level message since a certain time
func ErrorMsgSince(start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(nil, globalLogger.levels.Error, globalLogger.showErrorStack, start, v...)
}

// Output a ERROR-level message
func ErrorMsgf(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logErrorFormatted(nil, globalLogger.levels.Error, globalLogger.showErrorStack, format, v...)
}

// Output a FATAL-level message with information about the error
func Fatal(err error, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(err, globalLogger.levels.Fatal, globalLogger.showFatalStack, false, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output an FATAL log message since a certain time with information about the error
func FatalSince(err error, start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(err, globalLogger.levels.Fatal, globalLogger.showErrorStack, start, v...)
}

// Output a formatted FATAL-level message with information about the error
func Fatalf(err error, format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logErrorFormatted(err, globalLogger.levels.Fatal, globalLogger.showFatalStack, format, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a FATAL-level message
func FatalMsg(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(nil, globalLogger.levels.Fatal, globalLogger.showFatalStack, false, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output an FATAL-level message since a certain time
func FatalMsgSince(err error, start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(err, globalLogger.levels.Fatal, globalLogger.showErrorStack, start, v...)
}

// Output a formatted FATAL-level message
func FatalMsgf(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logErrorFormatted(nil, globalLogger.levels.Fatal, globalLogger.showFatalStack, format, v...)
	os.Exit(globalLogger.fatalExitCode)
}
