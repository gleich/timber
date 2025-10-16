package timber

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func format(level Level, v ...any) *strings.Builder {
	var out strings.Builder
	out.WriteString(time.Now().In(globalLogger.timezone).Format(globalLogger.timeFormat))
	out.WriteRune(' ')
	out.WriteString(level.renderedMsg)
	out.WriteRune(' ')

	for i, item := range v {
		if i > 0 {
			out.WriteRune(' ')
		}
		out.WriteString(fmt.Sprint(item))
	}
	return &out
}

func logNormal(level Level, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()
	globalLogger.normalOutput.logger.Print(format(level, v...).String())
}

func logError(err error, level Level, outputStack bool, v ...any) {
	globalLogger.mutex.RLock()
	defer globalLogger.mutex.RUnlock()

	out := format(level, v...)
	if err != nil {
		out.WriteRune('\n')
		out.WriteString(err.Error())
	}
	if outputStack {
		out.WriteRune('\n')
		stackTrace(out)
	}
	globalLogger.errOutput.logger.Print(out.String())
}

func stackTrace(builder *strings.Builder) {
	pcs := make([]uintptr, 64)
	n := runtime.Callers(3, pcs) // skip runtime.Callers, stackTrace, logError/Error*
	if n == 0 {
		return
	}
	pcs = pcs[:n]

	if len(pcs) < 2 {
		return
	}
	pcs = pcs[:len(pcs)-2]

	fr := runtime.CallersFrames(pcs)
	var frames []runtime.Frame
	for {
		f, more := fr.Next()
		frames = append(frames, f)
		if !more {
			break
		}
	}

	for i := 0; i < len(frames); i++ {
		if i+1 < len(frames) {
			fmt.Fprintf(
				builder,
				"#%d. %s() [%s:%d]\n",
				i+1,
				frames[i].Function,
				frames[i+1].File,
				frames[i+1].Line,
			)
		} else {
			fmt.Fprintf(builder, "#%d. %s()\n", i+1, frames[i].Function)
		}
	}
}

// Output a DEBUG log message
func Debug(v ...any) {
	logNormal(globalLogger.levels.Debug, v...)
}

// Output a DONE log message
func Done(v ...any) {
	logNormal(globalLogger.levels.Done, v...)
}

// Output an INFO log message
func Info(v ...any) {
	logNormal(globalLogger.levels.Info, v...)
}

// Output a WARN log message
func Warning(v ...any) {
	logNormal(globalLogger.levels.Warning, v...)
}

// Output an ERROR log message with information about the error
func Error(err error, v ...any) {
	logError(err, globalLogger.levels.Error, globalLogger.showErrorStack, v...)
}

// Output an ERROR log message
func ErrorMsg(v ...any) {
	logError(nil, globalLogger.levels.Error, globalLogger.showErrorStack, v...)
}

// Output a FATAL log message with information about the error
func Fatal(err error, v ...any) {
	logError(err, globalLogger.levels.Fatal, globalLogger.showFatalStack, v...)
	os.Exit(globalLogger.fatalExitCode)
}

// Output a FATAL log message
func FatalMsg(v ...any) {
	logError(nil, globalLogger.levels.Fatal, globalLogger.showFatalStack, v...)
	os.Exit(globalLogger.fatalExitCode)
}
