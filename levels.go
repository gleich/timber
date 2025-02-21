package timber

import "github.com/charmbracelet/lipgloss"

type Levels struct {
	Debug   Level
	Info    Level
	Done    Level
	Warning Level
	Error   Level
	Fatal   Level
}

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
	if normalLevels {
		levels := []*Level{
			&logger.levels.Debug,
			&logger.levels.Info,
			&logger.levels.Done,
			&logger.levels.Warning,
		}
		for _, level := range levels {
			renderLevel(level)
		}
	}
	if errLevels {
		levels := []*Level{&logger.levels.Error, &logger.levels.Fatal}
		for _, level := range levels {
			renderLevel(level)
		}
	}
}

func SetLevels(levels Levels) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels = levels
	renderLevels(globalLogger, true, true)
}

func SetDebug(l Level) {
	setLevel(&globalLogger.levels.Debug, l)
}

func SetDebugStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Debug, s)
}

func SetInfo(l Level) {
	setLevel(&globalLogger.levels.Info, l)
}

func SetInfoStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Info, s)
}

func SetDone(l Level) {
	setLevel(&globalLogger.levels.Done, l)
}

func SetDoneStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Done, s)
}

func SetWarning(l Level) {
	setLevel(&globalLogger.levels.Warning, l)
}

func SetWarningStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Warning, s)
}

func SetError(l Level) {
	setLevel(&globalLogger.levels.Error, l)
}

func SetErrorStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Error, s)
}

func SetFatal(l Level) {
	setLevel(&globalLogger.levels.Fatal, l)
}

func SetFatalStyle(s lipgloss.Style) {
	setLevelStyle(&globalLogger.levels.Fatal, s)
}
