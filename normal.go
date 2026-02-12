package timber

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
