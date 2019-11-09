package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launching")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
