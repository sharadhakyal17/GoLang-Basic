package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffe", "Tea", "Hot C"}
	fmt.Println(arr[1])

	arr[1] = "Green Tea"
	fmt.Println(arr[1])

	arr2 := arr
	fmt.Println(arr2)

	arr[1] = arr2[1]
	fmt.Println(arr[1])

	fmt.Println(arr, arr2)
}
