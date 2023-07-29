package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	// list item
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... %v\n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+b.tip)

	return fs

}

func (b *bill) updateTip(t float64) {
	b.tip = t
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
