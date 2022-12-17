package main // Package declaration. This is the entry point of the program. It allows us to use the main function and make it executable.

import "fmt"

func main() {
	x := 0
	for x < 5 {
		println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	for i, name := range names {
		fmt.Printf("The position at %v is %v\n", i, name)
	}
}
