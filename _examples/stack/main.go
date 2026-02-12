package main

import (
	"os"
	"sync"

	"go.mattglei.ch/timber"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	caller(wg)
	// go caller(wg)
	wg.Wait()

}

func caller(wg *sync.WaitGroup) {
	_, err := os.Stat("foo")
	if err != nil {
		timber.Fatal(err)
	}
	wg.Done()
}
