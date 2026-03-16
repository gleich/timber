package timber

import (
	"os"
	"time"
)

// Output an ERROR-level message with information about the error
func Error(err error, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(globalLogger.levels.Error, err, msg, vals, globalLogger.showErrorStack)
}

// Output an ERROR-level message since a certain time with information about the error
func ErrorSince(err error, start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(
		globalLogger.levels.Error,
		err,
		start,
		msg,
		vals,
		globalLogger.showErrorStack,
	)
}

// Output a ERROR-level message
func ErrorMsg(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(globalLogger.levels.Error, nil, msg, vals, globalLogger.showErrorStack)
}

// Output an ERROR-level message since a certain time
func ErrorMsgSince(start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(
		globalLogger.levels.Error,
		nil,
		start,
		msg,
		vals,
		globalLogger.showErrorStack,
	)
}

// Output a FATAL-level message with information about the error
func Fatal(err error, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(globalLogger.levels.Fatal, err, msg, vals, globalLogger.showFatalStack)
	os.Exit(globalLogger.fatalExitCode)
}

// Output an FATAL log message since a certain time with information about the error
func FatalSince(err error, start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(
		globalLogger.levels.Fatal,
		err,
		start,
		msg,
		vals,
		globalLogger.showErrorStack,
	)
}

// Output a FATAL-level message
func FatalMsg(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logError(globalLogger.levels.Fatal, nil, msg, vals, globalLogger.showFatalStack)
	os.Exit(globalLogger.fatalExitCode)
}

// Output an FATAL-level message since a certain time
func FatalMsgSince(err error, start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationError(
		globalLogger.levels.Fatal,
		err,
		start,
		msg,
		vals,
		globalLogger.showErrorStack,
	)
}
