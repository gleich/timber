package main

import (
	"fmt"
	"os"
	"time"

	"go.mattglei.ch/timber"
)

func main() {
	demos := []func(){
		func() {
			timber.Done("loaded up the program!")
			time.Sleep(2 * time.Second)
			timber.Done("waited 2 seconds")
		},
		func() {
			timber.Info("getting the current year")
			now := time.Now()
			timber.Info("current year is", now.Year())
		},
		func() {
			homeDir, _ := os.UserHomeDir()
			timber.Debug("user's home dir is", homeDir)
		},
		func() {
			now := time.Now()
			if now.Year() != 2004 {
				timber.Warning("current year isn't 2004")
			}
		},
		func() {
			timber.ErrorMsg("error message")
		},
		func() {
			fname := "invisible-file.txt"
			_, err := os.ReadFile(fname)
			if err != nil {
				timber.Error(err, "failed to read from", fname)
			}
		},
		// func() {
		// 	timber.FatalMsg("fatal message")
		// },
	}

	for _, demo := range demos {
		fmt.Println()
		demo()
		fmt.Println()
	}
}
