package main

import "fmt"

func changeMap(m map[string]float64) {
	m["pizza"] = 5.00
}

func main() {

	menu := map[string]float64{
		"pie":       5.95,
		"cake":      3.99,
		"ice cream": 4.95,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Printf("The menu item, %v is this price: %v \n", k, v)
	}

	// Maps are reference types. They are not passed by value. They are passed by reference. This means that if we change the value of a map inside of a function, the change will be reflected in the original map.

	changeMap(menu)
	fmt.Println(menu)
}
