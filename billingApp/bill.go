package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

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
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs

}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip // b.tip is a pointer to the tip value in the bill struct. We don't need to use the * operator to dereference the pointer because GO does it for us.
}

// Add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0664)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to a file")

}
