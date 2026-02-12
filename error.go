package timber

import "os"

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
