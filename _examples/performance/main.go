package main

import (
	"fmt"
	"log"
	"time"

	"pkg.mattglei.ch/timber"
)

func main() {
	timberResult := timeFunc(timber.Done)
	logResult := timeFunc(log.Println)

	fmt.Println()
	fmt.Println("timber:", timberResult)
	fmt.Println("log:", logResult)
}

func timeFunc(f func(v ...any)) string {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		f("foo", "bar")
	}
	return time.Since(start).String()
}
