package main

import (
	"go.mattglei.ch/timber"
)

func main() {
	timber.Structured(true)

	timber.Debug("hello world", timber.A("foo", "bar"))
}
