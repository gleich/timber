package timber

import "time"

// Output a DEBUG-level message
func Debug(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Debug, v...)
}

// Output a DEBUG-level message since a certain time
func DebugSince(start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Debug, start, v...)
}

// Output a formatted DEBUG-level message
func Debugf(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logFormatted(globalLogger.levels.Debug, format, v...)
}

// Output a DONE-level message
func Done(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Done, v...)
}

// Output a DONE-level message since a certain time
func DoneSince(start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, v...)
}

// Output a formatted DONE-level message
func Donef(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logFormatted(globalLogger.levels.Done, format, v...)
}

// Output a INFO-level message
func Info(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Info, v...)
}

// Output a INFO-level message since a certain time
func InfoSince(start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, v...)
}

// Output a formatted INFO-level message
func Infof(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logFormatted(globalLogger.levels.Info, format, v...)
}

// Output a WARN-level message
func Warning(v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Warning, v...)
}

// Output a WARNING-level message since a certain time
func WarningSince(start time.Time, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, v...)
}

// Output a formatted WARN-level message
func Warningf(format string, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logFormatted(globalLogger.levels.Warning, format, v...)
}
