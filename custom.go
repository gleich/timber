package timber

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var globalLogger *loggerOptions

type loggerOptions struct {
	mutex         sync.RWMutex
	out           io.Writer
	logger        *log.Logger
	errOut        io.Writer
	errLogger     *log.Logger
	renderer      *lipgloss.Renderer
	errRenderer   *lipgloss.Renderer
	extraOuts     []io.Writer
	extraErrOuts  []io.Writer
	fatalExitCode int
	showStack     bool
	timeFormat    string
	timezone      *time.Location
	levels        Levels
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
		l = loggerOptions{
			mutex:         sync.RWMutex{},
			out:           out,
			logger:        log.New(out, "", 0),
			errOut:        errOut,
			errLogger:     log.New(errOut, "", 0),
			renderer:      renderer,
			errRenderer:   errRenderer,
			extraOuts:     []io.Writer{},
			extraErrOuts:  []io.Writer{},
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

func updateNormalLogger() {
	globalLogger.out = io.MultiWriter(append(globalLogger.extraOuts, globalLogger.out)...)
}

func updateErrLogger() {
	globalLogger.errOut = io.MultiWriter(append(globalLogger.extraErrOuts, globalLogger.errOut)...)
}

// Set the output for Debug, Done, Warning, and Info.
//
// Default is os.Stdout
func Out(out io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.out = out
	globalLogger.renderer = lipgloss.NewRenderer(out)
	updateNormalLogger()
	renderLevels(globalLogger, true, false)
}

// Set the output for Fatal, FatalMsg, Error, and ErrorMsg.
//
// Default is os.Stderr
func ErrOut(out io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.errOut = out
	globalLogger.errRenderer = lipgloss.NewRenderer(out)
	updateErrLogger()
	renderLevels(globalLogger, false, true)
}

// Set the extra normal output destinations (e.g. logging to a file).
func ExtraOuts(outs []io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.extraOuts = outs
	updateNormalLogger()
}

// Set the extra error output destinations (e.g. logging to a file).
func ExtraErrOuts(outs []io.Writer) {
	globalLogger.mutex.Lock()
	defer globalLogger.mutex.Unlock()
	globalLogger.extraErrOuts = outs
	updateErrLogger()
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
