package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This will cause a panic because we are attempting to divide by zero.
	panic("Oh no! Something went wrong!")
}
