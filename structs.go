package main

import "fmt"

func main() {

	// Struct is a blueprints

	mybill := newBill("marios bill")

	mybill.addItem("Oinion", 12)
	mybill.addItem("Value", 2.3)
	mybill.addItem("Banana", 1.2)
	mybill.addItem("Orange", 3.3)
	mybill.addItem("Mango", 6.7)

	mybill.updateTip(10)

	fmt.Println(mybill.format())

}
