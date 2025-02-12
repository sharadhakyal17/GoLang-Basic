package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffe": {"Coffe", "Espresso", "Cappucino"},
		"tea":   {"Hot Tea", "Chai"},
	}

	fmt.Println(m)
	fmt.Println(m["coffe"][0])

	m["other"] = []string{"HChoc"}

	fmt.Println(m)

	delete(m, "coffe")
	fmt.Println(m)

	fmt.Println(m["tea"])

	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffe"] = []string{"Hot Tea"}

	fmt.Println(m2["coffe"])
	fmt.Println(m["coffe"])
}
