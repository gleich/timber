package main

import (
	"os"

	"go.mattglei.ch/timber"
)

func main() {
	_, err := os.Stat("foo")
	if err != nil {
		timber.Fatal(err, "failed to check status of foo")
	}
}
