package main

import "fmt"

func main() {
	//string
	var nameOne string = "Abebe"

	// automatically assign type by go
	var nameTwo = "Chale"

	// setting up a variable for a future use
	var nameThree string
	nameThree = "Hagos"

	// short hand variable declaration for only first time variable declaration and initialization
	// this kind of initialization cannot be used outside of a function
	nameFour := "Quantity"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Numbers: we have int and float types
	var ageOne int = 24
	var ageTwo = 34
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory | if you have an interest
	var numOne int8 = 25
	var numTwo int8 = -128
	//only posetive integers specified by u
	var numThree uint = 25

	fmt.Println(numOne, numTwo, numThree)

	// float must specify a bit size to be declared
	var scoreOne float32 = 2.3
	// using float64 has higher precision. We will be using float64 for most time
	var scoreTwo float64 = 3222223242.52242

	// dafult will be float64
	scoreThree := 2.7

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
