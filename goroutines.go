package main

import (
	"fmt"
	"time"
)

func main() {
	// Launch a goroutine to print "Hello" every 1 second.
	go func() {
		for {
			fmt.Println("Hello")
			time.Sleep(time.Second)
		}
	}()

	// Launch another goroutine to print "World" every 2 seconds.
	go func() {
		for {
			fmt.Println("World")
			time.Sleep(2 * time.Second)
		}
	}()

	// Sleep for 5 seconds to allow the goroutines to run.
	time.Sleep(5 * time.Second)
}
