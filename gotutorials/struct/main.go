package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Cipher's Bill")
	myBill.updateTip(5)
	myBill.addItem("Spaghetti", 10.99)
	fmt.Println(myBill.format())
}
