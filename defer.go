package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	/*
		By using defer, we can be sure that the Body is closed before the function returns.
		This way, we place it next to the code that opens the Body, and we don’t have to worry about closing it later.
		It still needs to be after the error check, though, because we don’t want to defer a function call if there’s an error.
		You might want to be carful about using defer in a loop, because it will be called after the loop is finished.
		If you open a lot of resources, you might run out of memory. In that case, you should close the resources as soon as you’re done with them.
		Finally, you could just delegate opening and closing the resource to another function, and use defer in that function.

		The order of execution is:
		1. The function is called
		2. The defer statements are executed in the reverse order
		3. The function returns
		4. Any panic is handled

	*/

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
