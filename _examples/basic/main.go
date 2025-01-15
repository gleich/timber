package main

import "pkg.mattglei.ch/timber"

func main() {
	msg := "Hello World!"
	timber.Debug(msg)
	timber.Done(msg)
	timber.Info(msg)
	timber.Warning(msg)
	timber.ErrorMsg(msg)
	timber.FatalMsg(msg)
}
