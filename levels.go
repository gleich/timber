package timber

import "github.com/charmbracelet/lipgloss"

// Levels used by timber for logging
type Levels struct {
	Debug   Level
	Info    Level
	Done    Level
	Warning Level
	Error   Level
	Fatal   Level
}

// A given level for logging
type Level struct {
	Style       lipgloss.Style
	Message     string
	renderedMsg string
}

func renderLevel(level *Level) {
	level.renderedMsg = level.Style.Render(level.Message)
}

func setLevelStyle(level *Level, style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	level.Style = style
	renderLevel(level)
}

func setLevel(currentLevel *Level, newLevel Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	*currentLevel = newLevel
	renderLevel(currentLevel)
}

func renderLevels(logger *logger, normalLevels bool, errLevels bool) {
	switch {
	case normalLevels:
		levels := []*Level{
			&logger.levels.Debug,
			&logger.levels.Info,
			&logger.levels.Done,
			&logger.levels.Warning,
		}
		for _, level := range levels {
			renderLevel(level)
		}
	case errLevels:
		levels := []*Level{&logger.levels.Error, &logger.levels.Fatal}
		for _, level := range levels {
			renderLevel(level)
		}
	}
}

// Set the levels that timber logs at.
//
// Default:
// DEBUG - Bold #2B95FF
// INFO  - Bold
// DONE  - Bold #30CE75
// WARN  - Bold #E1DC3F
// ERROR - Bold #FF4747
// FATAL - Bold #FF4747
func SetLevels(levels Levels) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels = levels
	renderLevels(globalLogger, true, true)
}

// Set the level for the debug level
func SetDebug(l Level) {
	setLevel(&globalLogger.levels.Debug, l)
}

// Set the style for the debug level
func SetDebugStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Debug, s)
}

// Set the level for the info level
func SetInfo(l Level) {
	setLevel(&globalLogger.levels.Info, l)
}

// Set the style for the info level
func SetInfoStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Info, s)
}

// Set the level for the done level
func SetDone(l Level) {
	setLevel(&globalLogger.levels.Done, l)
}

// Set the style for the done level
func SetDoneStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Done, s)
}

// Set the level for the warning level
func SetWarning(l Level) {
	setLevel(&globalLogger.levels.Warning, l)
}

// Set the style for the warning level
func SetWarningStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Warning, s)
}

// Set the level for the error level
func SetError(l Level) {
	setLevel(&globalLogger.levels.Error, l)
}

// Set the style for the error level
func SetErrorStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Error, s)
}

// Set the level for the fatal level
func SetFatal(l Level) {
	setLevel(&globalLogger.levels.Fatal, l)
}

// Set the style for the fatal level
func SetFatalStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Fatal, s)
}
