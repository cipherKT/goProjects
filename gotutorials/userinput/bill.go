package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   int
}

// method to make a new bill

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

//method to format the string
// here we are using the receiver function so this can be invoked by only bill object like b.format()
// cannot be invoked like format()

func (b Bill) format() string {
	fs := "Bill Breakdown : \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ....$%v \n", k+":", v)
		total += v
	}
	if b.tip != 0 {
		total += float64(b.tip)
	}
	fs += fmt.Sprintf("%-25v ...%0.2v \n", "Tip : ", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2v \n", "Total : ", total)

	return fs
}

// update tip
func (b *Bill) updateTip(new_tip int) {
	b.tip = new_tip
}

// add item
func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *Bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved")

}
