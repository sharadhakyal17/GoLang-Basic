package main

import "fmt"

func main() {

	s := "Hello, Gophers!"

	p := &s         //& is the sign for pointer reference
	fmt.Println(p)  // prints the memory address value
	fmt.Println(*p) // prints the actual value stored in the memory address, and called as de-referencing

	*p = "Bye, Gophers!" // This updates the actual value stored in s.
	fmt.Println(s)

	p = new(string) //new allocates the memory similar to pointer
	fmt.Println(p)

}
