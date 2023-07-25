package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello there, how are you?"

	fmt.Println(strings.Contains(greeting, "Hello"))

	// Doesn't chage the value only replace a new value
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	fmt.Println(strings.ToUpper(greeting))

	fmt.Println(strings.Index(greeting, "ll"))

	// Return a slice
	fmt.Println(strings.Split(greeting, " "))

}
