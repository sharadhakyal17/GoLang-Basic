package main

import (
	"fmt"
	"slices"
)

func main() {
	var arr []string
	fmt.Println(arr)

	arr = []string{"Coffe", "Tea", "Hot C"}
	fmt.Println(arr[1])

	arr = append(arr, "No", "Non-sense")

	arr[1] = "Green Tea"
	fmt.Println(arr[1])

	arr2 := arr
	fmt.Println(arr2)

	arr[1] = arr2[1]
	fmt.Println(arr[1])

	fmt.Println(arr, arr2)

	slices.Delete(arr, 1, 3)
	fmt.Println(arr)
}
