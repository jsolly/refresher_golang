package main // Package declaration. This is the entry point of the program. It allows us to use the main function and make it executable.

import "fmt"

func main() {
	var myArray [5]int = [5]int{1, 2, 3, 4, 5} // Array declaration. It is a fixed length list of items of the same type. It is a value type.
	var myArray2 = [5]int{1, 2, 3, 4, 5}       // Array declaration. It is a fixed length list of items of the same type. It is a value type.
	myArray3 := [5]int{1, 2, 3, 4, 5}          // Array declaration. It is a fixed length list of items of the same type. It is a value type.

	fmt.Println(len(myArray), len(myArray2), len(myArray3)) // Prints the length of the array

	// Slices
	mySlice := []int{1, 2, 3, 4, 5} // Slice declaration. It is a variable length list of items of the same type. There is no fixed length.
	fmt.Println(mySlice[1:3])       // Prints the items from index 1 to 3.

	// Appending
	mySlice = append(mySlice, 6, 7, 8) // Appends the items to the slice.

}
