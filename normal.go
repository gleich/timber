package timber

import "time"

// Output a DEBUG-level message
func Debug(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Debug, msg, vals)
}

// Output a DEBUG-level message since a certain time
func DebugSince(start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Debug, start, msg, vals)
}

// Output a DONE-level message
func Done(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Done, msg, vals)
}

// Output a DONE-level message since a certain time
func DoneSince(start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, vals)
}

// Output a INFO-level message
func Info(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Info, msg, vals)
}

// Output a INFO-level message since a certain time
func InfoSince(start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, vals)
}

// Output a WARN-level message
func Warning(msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logNormal(globalLogger.levels.Warning, msg, vals)
}

// Output a WARNING-level message since a certain time
func WarningSince(start time.Time, msg string, vals ...Value) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	logDurationNormal(globalLogger.levels.Done, start, msg, vals)
}
