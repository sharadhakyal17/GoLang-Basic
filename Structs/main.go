package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like to scream?")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	type menuItem struct {
		name string
		price map[string]float64
	}

	menu := []menuItem{
		{name:"coffee", prices: map[string]float64{"small":1.65, "medium":2.0, "large" : 3.0} },
		{name:"tea", prices: map[string]float64{"small":1.65, "medium":2.0, "large" : 3.0} }

	}

	fmt.Println(menu)



}