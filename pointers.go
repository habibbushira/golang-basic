package main

import "fmt"

// * mean get the value from the passed addres
func updateName(name *string) {
	*name = "Change"
}

func main() {
	name := "Original"

	fmt.Println("Memory address of name ", &name)

	m := &name
	fmt.Println("Copied memory address of name ", m)
	fmt.Println("Get the value from the memory address", *m)

	fmt.Println("Before", name)
	updateName(m)
	fmt.Println("After", name)
}
