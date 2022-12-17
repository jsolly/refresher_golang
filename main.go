package main

import "fmt"

// Import the foo package

func main() {
	// Call the foo() function from the foo package
	foo.foo()
	// Print a message to the command line
	fmt.Println("Hello from file1.go")
}
