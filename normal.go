package timber

import "time"

// Output a DEBUG-level message
func Debug(msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Debug, msg, attrs)
}

// Output a DEBUG-level message since a certain time
func DebugSince(start time.Time, msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Debug, start, msg, attrs)
}

// Output a DONE-level message
func Done(msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Done, msg, attrs)
}

// Output a DONE-level message since a certain time
func DoneSince(start time.Time, msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, attrs)
}

// Output a INFO-level message
func Info(msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Info, msg, attrs)
}

// Output a INFO-level message since a certain time
func InfoSince(start time.Time, msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Info, start, msg, attrs)
}

// Output a WARN-level message
func Warning(msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Warning, msg, attrs)
}

// Output a WARNING-level message since a certain time
func WarningSince(start time.Time, msg string, attrs ...Attr) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, attrs)
}
