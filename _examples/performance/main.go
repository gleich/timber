package main

import (
	"fmt"
	"log"
	"time"

	"go.mattglei.ch/timber"
)

func main() {
	logResult := timeFunc(log.Println)
	timberResult := timeFunc(timber.Done)

	fmt.Println()
	fmt.Println("timber:", timberResult)
	fmt.Println("log:", logResult)
}

func timeFunc(f func(v ...any)) string {
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		f("foo", "bar")
	}
	return time.Since(start).String()
}
