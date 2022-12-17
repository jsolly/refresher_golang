package main

import (
	"fmt"
	"math"
	"strings"
)

func sayHello(n string) {
	fmt.Printf("Hello %s \n", n)
}

func cyclenames(n []string, f func(string)) {
	for _, n := range n {
		f(n)
	}
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	sayHello("John")
	names := []string{"John", "Paul", "George", "Ringo"}
	cyclenames(names, sayHello)

	a1 := circleArea(10.0)
	fmt.Printf("Area of circle with radius 10 is %0.1f \n", a1)

	// Multiple return values
	fname, lname := getInitials("John Paul")
	fmt.Printf("First name: %s, Last name: %s \n", fname, lname)

}
