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

func (l *Level) render() {
	l.renderedMsg = l.Style.Render(l.Message)
}

func (l *Level) style(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	l.Style = style
	l.render()
}

func (l *Level) set(newLevel Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	*l = newLevel
	l.render()
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
			level.render()
		}
	case errLevels:
		levels := []*Level{&logger.levels.Error, &logger.levels.Fatal}
		for _, level := range levels {
			level.render()
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
	globalLogger.levels.Debug.set(l)
}

// Set the style for the debug level
func SetDebugStyle(s lipgloss.Style) {
	globalLogger.levels.Debug.style(s)
}

// Set the level for the info level
func SetInfo(l Level) {
	globalLogger.levels.Info.set(l)
}

// Set the style for the info level
func SetInfoStyle(s lipgloss.Style) {
	globalLogger.levels.Info.style(s)
}

// Set the level for the done level
func SetDone(l Level) {
	globalLogger.levels.Done.set(l)
}

// Set the style for the done level
func SetDoneStyle(s lipgloss.Style) {
	globalLogger.levels.Done.style(s)
}

// Set the level for the warning level
func SetWarning(l Level) {
	globalLogger.levels.Warning.set(l)
}

// Set the style for the warning level
func SetWarningStyle(s lipgloss.Style) {
	globalLogger.levels.Warning.style(s)
}

// Set the level for the error level
func SetError(l Level) {
	globalLogger.levels.Error.set(l)
}

// Set the style for the error level
func SetErrorStyle(s lipgloss.Style) {
	globalLogger.levels.Error.style(s)
}

// Set the level for the fatal level
func SetFatal(l Level) {
	globalLogger.levels.Fatal.set(l)
}

// Set the style for the fatal level
func SetFatalStyle(s lipgloss.Style) {
	globalLogger.levels.Fatal.style(s)
}
