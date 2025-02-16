package timber

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var globalLogger *logger

type logger struct {
	mutex         sync.RWMutex
	normalOutput  output
	errOutput     output
	fatalExitCode int
	showStack     bool
	timeFormat    string
	timezone      *time.Location
	levels        Levels
}

type output struct {
	logger   *log.Logger
	renderer *lipgloss.Renderer
	writer   io.Writer
}

func init() {
	var (
		out         = os.Stdout
		errOut      = os.Stderr
		renderer    = lipgloss.NewRenderer(out)
		errRenderer = lipgloss.NewRenderer(errOut)
		bold        = renderer.NewStyle().Bold(true)
		errBold     = errRenderer.NewStyle().Bold(true)
		errStyle    = errRenderer.NewStyle().Inherit(errBold).
				Foreground(lipgloss.Color("#FF4747"))
		l = logger{
			mutex: sync.RWMutex{},
			normalOutput: output{
				logger:   log.New(out, "", 0),
				writer:   out,
				renderer: renderer,
			},
			errOutput: output{
				logger:   log.New(errOut, "", 0),
				writer:   errOut,
				renderer: errRenderer,
			},
			fatalExitCode: 1,
			showStack:     true,
			timeFormat:    "01/02/2006 15:04:05 MST",
			timezone:      time.UTC,
			levels: Levels{
				Debug: Level{
					Message: "DEBUG",
					Style: renderer.NewStyle().
						Inherit(bold).
						Foreground(lipgloss.Color("#2B95FF")),
				},
				Info: Level{
					Message: "INFO ",
					Style:   bold,
				},
				Done: Level{
					Message: "DONE ",
					Style: renderer.NewStyle().
						Inherit(bold).
						Foreground(lipgloss.Color("#30CE75")),
				},
				Warning: Level{
					Message: "WARN ",
					Style: renderer.NewStyle().
						Inherit(bold).
						Foreground(lipgloss.Color("#E1DC3F")),
				},
				Error: Level{
					Message: "ERROR",
					Style:   errStyle,
				},
				Fatal: Level{
					Message: "FATAL",
					Style:   errStyle,
				},
			},
		}
	)
	renderLevels(&l, true, true)
	globalLogger = &l
}

// Set the output for Debug, Done, Warning, and Info.
//
// Default is os.Stdout
func Out(writer io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.normalOutput.writer = writer
	globalLogger.normalOutput.renderer = lipgloss.NewRenderer(writer)
	renderLevels(globalLogger, true, false)
}

// Set the output for Fatal, FatalMsg, Error, and ErrorMsg.
//
// Default is os.Stderr
func ErrOut(writer io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.errOutput.writer = writer
	globalLogger.errOutput.renderer = lipgloss.NewRenderer(writer)
	renderLevels(globalLogger, false, true)
}

// Set the exit code used by Fatal and FatalMsg.
//
// Default is 1
func FatalExitCode(code int) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.fatalExitCode = code
}

// Set if the stack trace should be shown or not when calling Error or Fatal.
//
// Default is true
func ShowStack(show bool) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.showStack = show
}

// Set the time format that timestamps are formatted with.
//
// Default is 01/02/2006 15:04:05 MST
func TimeFormat(format string) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.timeFormat = format
}

// Set the timezone that timestamps are logged in.
//
// Default is time.UTC
func Timezone(loc *time.Location) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.timezone = loc
}
