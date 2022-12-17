package main

import "fmt"

func main() {
	age := 44
	fmt.Println(age > 50)  // Prints false
	fmt.Println(age < 50)  // Prints true
	fmt.Println(age == 50) // Prints false
	fmt.Println(age != 50) // Prints true

	if age > 50 {
		fmt.Println("You are old")
	} else if age < 30 {
		fmt.Println("You are young")
	} else {
		fmt.Println("You are middle aged")
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for i, name := range names {
		if i == 1 {
			continue // Skips the iteration
		}
		if i > 2 {
			break // Breaks the loop
		}
		fmt.Println(name)
	}
}
