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

func SetDebug(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Debug = level
	renderLevel(&globalLogger.levels.Debug)
}

func SetDebugStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Debug.Style = style
	renderLevel(&globalLogger.levels.Debug)
}

func SetInfo(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Info = level
	renderLevel(&globalLogger.levels.Info)
}

func SetInfoStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Info.Style = style
	renderLevel(&globalLogger.levels.Info)
}

func SetDone(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Done = level
	renderLevel(&globalLogger.levels.Done)
}

func SetDoneStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Done.Style = style
	renderLevel(&globalLogger.levels.Done)
}

func SetWarning(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Warning = level
	renderLevel(&globalLogger.levels.Warning)
}

func SetWarningStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Warning.Style = style
	renderLevel(&globalLogger.levels.Warning)
}

func SetError(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Error = level
	renderLevel(&globalLogger.levels.Error)
}

func SetErrorStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Error.Style = style
	renderLevel(&globalLogger.levels.Error)
}

func SetFatal(level Level) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Fatal = level
	renderLevel(&globalLogger.levels.Fatal)
}

func SetFatalStyle(style lipgloss.Style) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.levels.Fatal.Style = style
	renderLevel(&globalLogger.levels.Fatal)
}
