package main

import (
	"os"

	"go.mattglei.ch/timber"
)

func main() {
	_, err := os.Stat("sample.txt")
	if err != nil {
		timber.Fatal(err, "failed to read from file")
	}
}
