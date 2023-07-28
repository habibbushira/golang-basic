package main

import "fmt"

// Pointer wrapper values: Slices, Maps, Functions

// Non-pointer values: String, Int, Float, Boolean, Array, Struct

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 2.33,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as keytype

	phonebook := map[int]string{
		2324242: "Test",
		3420224: "Value",
		2223213: "Ional",
	}

	phonebook[3420224] = "New"
	fmt.Println(phonebook)
}
