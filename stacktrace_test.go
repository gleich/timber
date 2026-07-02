package timber

import (
	"runtime"
	"strings"
	"sync"
	"testing"
)

func requireNoEmpty(t *testing.T, frames []frame) {
	t.Helper()
	for i, f := range frames {
		if strings.TrimSpace(f.function) == "" {
			t.Fatalf("frame[%d] has empty function: %#v", i, f)
		}
	}
}

func requireFunctionsEndWithParens(t *testing.T, frames []frame) {
	t.Helper()
	for i, f := range frames {
		if !strings.HasSuffix(f.function, "()") {
			t.Fatalf("frame[%d] function does not end with (): %q", i, f.function)
		}
		if strings.Contains(f.function, "{0x") || strings.Contains(f.function, "0x") {
			t.Fatalf("frame[%d] function still contains hex/params: %q", i, f.function)
		}
	}
}

func requirePathLooksLikeFileLine(t *testing.T, frames []frame) {
	t.Helper()
	for i, f := range frames {
		if !strings.Contains(f.path, ".go:") && f.path != "" {
			t.Fatalf("frame[%d] path does not contain .go:: %q", i, f.path)
		}
		col := strings.LastIndexByte(f.path, ':')
		if (col == -1 || col == len(f.path)-1) && f.path != "" {
			t.Fatalf("frame[%d] path missing line suffix: %q", i, f.path)
		}
		for _, r := range f.path[col+1:] {
			if r < '0' || r > '9' {
				t.Fatalf("frame[%d] path line suffix not numeric: %q", i, f.path)
			}
		}
	}
}

func TestFramesFromCall_Stress_Concurrent(t *testing.T) {
	workers := max(4, runtime.GOMAXPROCS(0)*4)
	itersPerWorker := 500

	var wg sync.WaitGroup
	wg.Add(workers)

	errCh := make(chan string, workers*itersPerWorker)

	for range workers {
		go func() {
			defer wg.Done()

			for i := range itersPerWorker {
				skip := i % 6

				var frames []frame
				func() {
					defer func() {
						if r := recover(); r != nil {
							errCh <- "panic: " + stringifyRecover(r)
						}
					}()
					frames = framesFromCall(skip)
				}()

				if frames == nil {
					errCh <- "frames is nil"
					continue
				}

				if len(frames) == 0 {
					continue
				}

				func() {
					defer func() {
						if r := recover(); r != nil {
							errCh <- "assertion panic: " + stringifyRecover(r)
						}
					}()

					for _, f := range frames {
						if !(strings.HasSuffix(f.function, "()") || strings.Contains(f.function, "in goroutine")) {
							errCh <- "function not normalized: " + f.function
							return
						}
						if strings.Contains(f.function, "0x") {
							errCh <- "function contains params/hex: " + f.function
							return
						}
						if strings.Contains(f.path, "+0x") {
							errCh <- "path still contains +0x: " + f.path
							return
						}
						if !(strings.Contains(f.path, ".go:") || f.path == "") {
							errCh <- "path does not look like go file: " + f.path
							return
						}
					}
				}()
			}
		}()
	}

	wg.Wait()
	close(errCh)

	for msg := range errCh {
		t.Fatalf("framesFromCall stress failure: %s", msg)
	}
}

func TestFramesFromCall_BasicInvariants(t *testing.T) {
	frames := framesFromCall(0)
	if len(frames) == 0 {
		t.Fatalf("expected some frames, got 0")
	}

	requireNoEmpty(t, frames)
	requireFunctionsEndWithParens(t, frames)

	requirePathLooksLikeFileLine(t, frames)

	found := false
	for _, f := range frames {
		if strings.Contains(f.function, "TestFramesFromCall_BasicInvariants") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected to find this test function in stack frames, but did not")
	}
}

func stringifyRecover(r any) string {
	switch v := r.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return "unknown"
	}
}
