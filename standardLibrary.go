package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))          // Returns true if the string contains the substring
	fmt.Println(strings.Count("test", "t"))              // Returns the number of times the substring appears in the string
	fmt.Println(strings.HasPrefix("test", "te"))         // Returns true if the string starts with the substring
	fmt.Println(strings.HasSuffix("test", "st"))         // Returns true if the string ends with the substring
	fmt.Println(strings.Index("test", "e"))              // Returns the index of the first instance of the substring
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))   // Joins the elements of the slice to a string using the separator
	fmt.Println(strings.Repeat("a", 5))                  // Repeats the string 5 times
	fmt.Println(strings.Replace("foo", "o", "0", -1))    // Replaces all instances of the substring with the new substring
	fmt.Println(strings.ToUpper("test"))                 // Converts the string to uppercase
	fmt.Println(strings.ToLower("TEST"))                 // Converts the string to lowercase
	fmt.Println(strings.Split("a-b-c-d-e", "-"))         // Splits the string into a slice of substrings
	fmt.Println(strings.Trim(" !!! Achtung !!! ", "! ")) // Removes the leading and trailing characters

	ages := []int{35, 15, 25}
	sort.Ints(ages) // Sorts the slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 15) // Returns the index where the item would be inserted
	fmt.Println(index)

	names := []string{"John", "Paul", "George", "Ringo"}
	sort.Strings(names)
	fmt.Println(sort.SearchStrings(names, "John"))

}
