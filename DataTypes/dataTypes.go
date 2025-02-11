package main

import "fmt"

func main() {

	var a int
	a = 10
	fmt.Println(a)

	var b int = 20
	fmt.Println(b)

	var c = 30
	fmt.Println(c)

	var g = true
	fmt.Println(g)

	d := 40
	fmt.Println(d)

	//type conversion
	f_value := 3.14
	fmt.Println(f_value)

	var convert_f_value_int int = int(f_value)
	fmt.Println(convert_f_value_int)

}
