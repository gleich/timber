package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.mattglei.ch/timber"
)

func main() {
	demos := []func(time.Time){
		func(start time.Time) {
			timber.DoneSince(start, "hello")
		},
		func(start time.Time) {
			now := time.Now()
			timber.InfoSince(start, "got current year", timber.A("year", now.Year()))
		},
		func(start time.Time) {
			homeDir, _ := os.UserHomeDir()
			timber.DebugSince(start, "got user's home directory", timber.A("path", homeDir))
		},
		func(start time.Time) {
			now := time.Now()
			if now.Year() != 2004 {
				timber.WarningSince(start, "current year isn't 2004")
			}
		},
		func(start time.Time) {
			timber.ErrorMsgSince(start, "error message")
		},
		func(start time.Time) {
			fname := "invisible-file.txt"
			_, err := os.ReadFile(fname)
			if err != nil {
				timber.ErrorSince(err, start, "failed to read file", timber.A("filename", fname))
			}
		},
	}

	for _, demo := range demos {
		fmt.Println()
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
		demo(start)
		fmt.Println()
	}
}
