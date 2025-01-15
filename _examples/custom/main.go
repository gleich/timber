package main

import (
	"math/rand"
	"time"

	"pkg.mattglei.ch/timber"
)

func main() {
	timber.SetTimezone(time.Local)
	timber.SetTimeFormat("Mon Jan 2 15:04:05 MST 2006")
	timber.SetFatalExitCode(0)

	randCap := 100
	timber.Debug(rand.Intn(randCap))
	timber.Info(rand.Intn(randCap))
	timber.Done(rand.Intn(randCap))
	timber.Warning(rand.Intn(randCap))
	timber.ErrorMsg(rand.Intn(randCap))
	timber.FatalMsg(rand.Intn(randCap))
}
