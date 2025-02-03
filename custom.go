package timber

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var logger loggerOptions = *defaultLogger()

type loggerOptions struct {
	mutex           *sync.RWMutex
	normalLogger    *log.Logger
	errLogger       *log.Logger
	normalOut       io.Writer
	errOut          io.Writer
	normalRenderer  *lipgloss.Renderer
	errRenderer     *lipgloss.Renderer
	extraNormalOuts []io.Writer
	extraErrOuts    []io.Writer
	fatalExitCode   int
	showStack       bool
	timeFormat      string
	timezone        *time.Location
	colors          Colors
}

type Colors struct {
	DebugStyle   lipgloss.Style
	InfoStyle    lipgloss.Style
	DoneStyle    lipgloss.Style
	WarningStyle lipgloss.Style
	ErrorStyle   lipgloss.Style
	FatalStyle   lipgloss.Style
}

const (
	defaultDebugColor = "#2B95FF"
	defaultDoneColor  = "#30CE75"
	defaultWarnColor  = "#E1DC3F"
	defaultFatalColor = "#FF4747"
	defaultErrorColor = "#FF4747"
)

// Initialize the default logger with the default values
func defaultLogger() *loggerOptions {
	normalOut := os.Stdout
	errOut := os.Stderr
	l := loggerOptions{
		mutex:           &sync.RWMutex{},
		normalLogger:    log.New(normalOut, "", 0),
		errLogger:       log.New(errOut, "", 0),
		normalOut:       normalOut,
		errOut:          errOut,
		normalRenderer:  lipgloss.NewRenderer(normalOut),
		errRenderer:     lipgloss.NewRenderer(errOut),
		extraNormalOuts: []io.Writer{},
		extraErrOuts:    []io.Writer{},
		fatalExitCode:   1,
		showStack:       true,
		timeFormat:      "01/02/2006 15:04:05 MST",
		timezone:        time.UTC,
	}
	l.colors = Colors{
		DebugStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(defaultDebugColor)).
			Bold(true),
		InfoStyle: l.normalRenderer.NewStyle().Bold(true),
		WarningStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(defaultWarnColor)).
			Bold(true),
		DoneStyle: l.normalRenderer.NewStyle().
			Foreground(lipgloss.Color(defaultDoneColor)).
			Bold(true),
		FatalStyle: l.errRenderer.NewStyle().
			Foreground(lipgloss.Color(defaultFatalColor)).
			Bold(true),
		ErrorStyle: l.errRenderer.NewStyle().
			Foreground(lipgloss.Color(defaultErrorColor)).
			Bold(true),
	}
	return &l
}

func updateNormalLogger() {
	logger.normalLogger = log.New(
		io.MultiWriter(append(logger.extraNormalOuts, logger.normalOut)...),
		"",
		0,
	)
}

func updateErrLogger() {
	logger.errLogger = log.New(
		io.MultiWriter(append(logger.extraErrOuts, logger.errOut)...),
		"",
		0,
	)
}

// Set the output or Debug, Done, Warning, and Info.
//
// Default is os.Stdout
func SetNormalOut(out io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.normalOut = out
	logger.normalRenderer = lipgloss.NewRenderer(out)
	updateNormalLogger()
	// rerendering colors incase new output doesn't support colors
	logger.colors.DebugStyle = logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(defaultDebugColor)).
		Bold(true)
	logger.colors.InfoStyle = logger.normalRenderer.NewStyle().Bold(true)
	logger.colors.DoneStyle = logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(defaultDoneColor)).
		Bold(true)
	logger.colors.WarningStyle = logger.normalRenderer.NewStyle().
		Foreground(lipgloss.Color(defaultWarnColor)).
		Bold(true)
}

// Set the output or Fatal, FatalMsg, Error, and ErrorMsg.
//
// Default is os.Stderr
func SetErrOut(out io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.errOut = out
	logger.errRenderer = lipgloss.NewRenderer(out)
	updateErrLogger()
	// rerendering colors incase new output doesn't support colors
	logger.colors.ErrorStyle = logger.errRenderer.NewStyle().
		Foreground(lipgloss.Color(defaultErrorColor)).
		Bold(true)
	logger.colors.FatalStyle = logger.errRenderer.NewStyle().
		Foreground(lipgloss.Color(defaultFatalColor)).
		Bold(true)
}

// Set the colors used for logging
func SetColors(colors Colors) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.colors.DebugStyle = logger.normalRenderer.NewStyle().Inherit(colors.DebugStyle)
	logger.colors.InfoStyle = logger.normalRenderer.NewStyle().Inherit(colors.InfoStyle)
	logger.colors.DoneStyle = logger.normalRenderer.NewStyle().Inherit(colors.DoneStyle)
	logger.colors.WarningStyle = logger.normalRenderer.NewStyle().Inherit(colors.WarningStyle)
	logger.colors.ErrorStyle = logger.errRenderer.NewStyle().Inherit(colors.ErrorStyle)
	logger.colors.FatalStyle = logger.errRenderer.NewStyle().Inherit(colors.FatalStyle)
}

// Set the extra normal out destinations (e.g. logging to a file)
func SetExtraNormalOuts(outs []io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.extraNormalOuts = outs
	updateNormalLogger()
}

// Set the extra normal out destinations (e.g. logging to a file)
func SetExtraErrOuts(outs []io.Writer) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.extraErrOuts = outs
	updateErrLogger()
}

// Set the exit code used by Fatal and FatalMsg.
//
// Default is 1
func SetFatalExitCode(code int) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.fatalExitCode = code
}

// Set if the stack trace should be shown or not when calling Error or Fatal.
//
// Default is true
func SetShowStack(show bool) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.showStack = show
}

// Set the time format
//
// Default is 01/02/2006 15:04:05 MST
func SetTimeFormat(format string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.timeFormat = format
}

// Set the timezone
//
// Default is time.UTC
func SetTimezone(loc *time.Location) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.timezone = loc
}
