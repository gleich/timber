package main

import (
	"os"

	"pkg.mattglei.ch/timber"
)

func main() {
	fname := "sample.txt"

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		timber.Warning(fname, "doesn't exist. Creating now")
		err := os.WriteFile(fname, []byte("testing"), 0655)
		if err != nil {
			timber.Fatal(err, "Failed to write to", fname)
		}
		timber.Done("Wrote to", fname)
	} else {
		timber.Info("Reading from file")
		b, err := os.ReadFile(fname)
		if err != nil {
			timber.Fatal(err, "Failed to read from", fname)
		}
		timber.Done("Read from", fname, "with content of", string(b))
	}
}
