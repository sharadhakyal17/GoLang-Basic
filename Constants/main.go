package main

import "fmt"

func main() {
	const a = 42
	var i int = a
	var f float64 = a

	fmt.Println(i, f)

	//this fails

	// const d float64 = 3 // explicitly specifying const variable type will fail and only with the same type conversion will work.
	// var g int = d
	// var h float32 = d

	// fmt.Println(g, h)
}
