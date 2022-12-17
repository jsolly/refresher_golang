import "fmt"

func foo() {
	fmt.Println("Hello, World!")

	var nameOne string = "John Solly"
	var nameTwo string
	var nameThree = "Holiday"
	nameFour := "Bob" // short hand declaration. Can only be used inside a function.

	// Bits and Memory
	// 1 byte = 8 bits
	var age int8 = 40            // int8 has a range of -128 to 127. If we try to assign a value greater than 127, it will throw an error.
	var ageTwo uint8 = 40        // uint8 has a range of 0 to 255. If we try to assign a value greater than 255, it will throw an error. It is used for positive numbers.
	var scoreOne float32 = 25.99 // used for decimals.
	fmt.Println(nameOne, nameTwo, nameThree, nameFour, age, ageTwo, scoreOne)
	fmt.Printf("My name is %v and I am %v years old. \n", nameOne, age) // Using Printf to print with a format specifier. %v is a placeholder for a value.
	fmt.Printf("My name is %q and I am %v years old. \n", nameOne, age) // Using Printf to print with a format specifier. %q is a placeholder for a value and it will be printed with double quotes.
	fmt.Printf("My age is of type %T \n", age)                          // Using Printf to print with a format specifier. %T is a placeholder for a type.
	fmt.Printf("My score is %0.1f \n", 25.99)                           // Using Printf to print with a format specifier. %0.1f is a placeholder for a value and it will be printed with 1 decimal place.
	// Save formatted strings
	var myString string = fmt.Sprintf("My name is %v and I am %v years old. \n", nameOne, age)
	fmt.Println(myString)
}
