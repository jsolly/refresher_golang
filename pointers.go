package main

import "fmt"

// Pointers
func updateName(x *string) { // *string is a pointer to a string
	*x = "Bobby" // *x is the value of the variable x. We are dereferencing the pointer to get the value of the variable x. We are then updating the value of the variable x.
}

func main() {
	name := "Sam"
	fmt.Println("Memory address of name:", &name)
	name_pointer := &name                        // &name is the memory address of the variable name.
	fmt.Println("value of name:", *name_pointer) // *name_pointer is the value of the variable name. We are dereferencing the pointer to get the value of the variable name.
	fmt.Println((name == *name_pointer))         // true because the value of the variable name is the same as the value of the variable name_pointer.
	fmt.Println(name)
	updateName(&name) // &name is the address of the variable name. We are passing the address of the variable name to the function updateName.
	fmt.Println(name)

}
