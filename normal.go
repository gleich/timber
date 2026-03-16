package timber

import "time"

// Output a DEBUG-level message
func Debug(msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Debug, msg, attributes)
}

// Output a DEBUG-level message since a certain time
func DebugSince(start time.Time, msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Debug, start, msg, attributes)
}

// Output a DONE-level message
func Done(msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Done, msg, attributes)
}

// Output a DONE-level message since a certain time
func DoneSince(start time.Time, msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, attributes)
}

// Output a INFO-level message
func Info(msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Info, msg, attributes)
}

// Output a INFO-level message since a certain time
func InfoSince(start time.Time, msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, attributes)
}

// Output a WARN-level message
func Warning(msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Warning, msg, attributes)
}

// Output a WARNING-level message since a certain time
func WarningSince(start time.Time, msg string, attributes ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, attributes)
}
