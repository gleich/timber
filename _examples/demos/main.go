package main

import (
	"fmt"
	"os"
	"time"

	"pkg.mattglei.ch/timber"
)

func main() {
	demos := []func(){
		func() {
			timber.Done("Loaded up the program!")
			time.Sleep(2 * time.Second)
			timber.Done("Waited 2 seconds")
		},
		func() {
			timber.Info("Getting the current year")
			now := time.Now()
			timber.Info("Current year is", now.Year())
		},
		func() {
			homeDir, _ := os.UserHomeDir()
			timber.Debug("User's home dir is", homeDir)
		},
		func() {
			now := time.Now()
			if now.Year() != 2004 {
				timber.Warning("Current year isn't 2004")
			}
		},
		func() {
			fname := "invisible-file.txt"
			_, err := os.ReadFile(fname)
			if err != nil {
				timber.Error(err, "Failed to read from", fname)
			}
		},
		func() {
			timber.FatalMsg("Fatal message")
		},
	}

	for _, demo := range demos {
		fmt.Println()
		demo()
		fmt.Println()
	}
}
