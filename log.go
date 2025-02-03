package timber

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/errors"
)

type logLevel string

const (
	debugLevel   logLevel = "DEBUG"
	doneLevel    logLevel = "DONE "
	infoLevel    logLevel = "INFO "
	warningLevel logLevel = "WARN "
	errorLevel   logLevel = "ERROR"
	fatalLevel   logLevel = "FATAL"
)

func format(level logLevel, color lipgloss.Style, v ...any) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s ", time.Now().In(logger.timezone).Format(logger.timeFormat))
	fmt.Fprintf(&b, "%s ", color.Render(string(level)))
	for i, item := range v {
		if i > 0 {
			b.WriteString(" ")
		}
		fmt.Fprint(&b, item)
	}
	return b.String()
}

// Normal log output
func logNormal(level logLevel, color lipgloss.Style, v ...any) {
	logger.mutex.RLock()
	defer logger.mutex.RUnlock()
	logger.normalLogger.Println(format(level, color, v...))
}

func logError(err error, level logLevel, color lipgloss.Style, v ...any) {
	logger.mutex.RLock()
	defer logger.mutex.RUnlock()
	out := format(level, color, v...)
	if err != nil && logger.showStack {
		out += fmt.Sprintf("\n%+v", errors.WithStack(err))
	} else if err != nil {

		out += fmt.Sprintf("\n%s", err)
	}
	logger.errLogger.Println(out)
}

// Output a INFO log message
func Debug(v ...any) {
	logNormal(debugLevel, logger.colors.DebugStyle, v...)
}

// Output a DONE log message
func Done(v ...any) {
	logNormal(doneLevel, logger.colors.DoneStyle, v...)
}

// Output a INFO log message
func Info(v ...any) {
	logNormal(infoLevel, logger.colors.InfoStyle, v...)
}

// Output a WARN log message
func Warning(v ...any) {
	logNormal(warningLevel, logger.colors.WarningStyle, v...)
}

// Output a ERROR log message with information about the error
func Error(err error, v ...any) {
	logError(err, errorLevel, logger.colors.ErrorStyle, v...)
}

// Output a ERROR log message
func ErrorMsg(v ...any) {
	logError(nil, errorLevel, logger.colors.ErrorStyle, v...)
}

// Output a FATAL log message with information about the error
func Fatal(err error, v ...any) {
	logError(err, fatalLevel, logger.colors.FatalStyle, v...)
	os.Exit(logger.fatalExitCode)
}

// Output a FATAL log message
func FatalMsg(v ...any) {
	logError(nil, fatalLevel, logger.colors.FatalStyle, v...)
	os.Exit(logger.fatalExitCode)
}
