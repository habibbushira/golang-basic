package main

import "fmt"

func main() {
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("newLine\n")

	//To add new line automatically
	fmt.Println("Hi there,")

	//To pring variables in fmt
	age := 13
	name := "Ousman"
	fmt.Println("Age: ", age, "Name: ", name)

	// Printf(formated string) %_ is a format specifiers
	//%v for default variables
	fmt.Printf("my age is %v and my name is %v\n", age, name)

	//%q to quote the variable
	fmt.Printf("my age is %v and my name is %q\n", age, name)

	//%T to output the variable type
	fmt.Printf("age type is %T\n", age)

	//%f to output a float value
	fmt.Printf("Float is %f\n", 42.22)

	//%_f to limit the number of decimals in the output
	fmt.Printf("Specified float is %0.1f\n", 42.22)

	//Sprintf(save formated string)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("SAved string, ", str)
}
