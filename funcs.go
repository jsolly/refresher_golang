package main

import (
	"fmt"
	"math"
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

func main() {
	sayHello("John")
	names := []string{"John", "Paul", "George", "Ringo"}
	cyclenames(names, sayHello)

	a1 := circleArea(10.0)
	fmt.Printf("Area of circle with radius 10 is %0.1f \n", a1)
}
