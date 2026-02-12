package timber

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type frame struct {
	function string
	path     string
}

func framesFromCall(skip int) []frame {
	frames := []frame{}
	out := string(debug.Stack())
	lines := strings.Split(out, "\n")

	// skip first line as that is information about the current goroutine that we do not care about
	for i := 1 + (skip * 2); i < len(lines)-1; i++ {
		f := frame{}
		function := lines[i]
		parameterStart := strings.IndexRune(function, '(')
		if parameterStart != -1 {
			function = fmt.Sprintf("%s()", function[:parameterStart])
		}
		f.function = strings.TrimPrefix(
			strings.TrimSuffix(function, " in goroutine"),
			"created by ",
		)

		pathParts := strings.Split(lines[i+1], " ")
		f.path = strings.TrimSpace(strings.Join(pathParts[:len(pathParts)-1], " "))
		frames = append(frames, f)
		i++ // increase twice because we just processed 2 lines
	}

	if len(frames) <= 1 {
		return frames
	}
	// path for each frame should be shifted up 1
	for i := 1; i < len(frames); i++ {
		frames[i-1].path = frames[i].path
	}
	frames[len(frames)-1].path = ""

	return frames
}

func stackTrace(builder *strings.Builder) {
	frames := framesFromCall(4)
	for i, f := range frames {
		trace := fmt.Sprintf("%d. %s", i+1, f.function)
		if f.path != "" {
			trace = fmt.Sprintf("%s %s", trace, globalLogger.stackPathStyle.Render(
				fmt.Sprintf("[%s]", f.path),
			))
		}
		fmt.Fprintln(builder, trace)
	}
}
