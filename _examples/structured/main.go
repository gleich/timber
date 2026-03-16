package main

import (
	"os"

	"go.mattglei.ch/timber"
)

func main() {
	// timber.Structured(true)
	// start := time.Now()

	// time.Sleep(2 * time.Second)

	_, err := os.ReadFile("hello")
	timber.Fatal(err, "nooo this didn't work", timber.V{"foo", 10}, timber.V{"remove", true})
}
