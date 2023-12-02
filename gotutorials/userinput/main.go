package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the name of the Bill : ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Enter the name of the Bill : ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a - add item , s - save bill , t - add tip) : ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name : ", reader)
		price, _ := getInput("Item price : ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number : ")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount ($) : ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number : ")
			promptOptions(b)
		}
		intTip := int(t)
		b.updateTip(intTip)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("that was not a valid option ...")
		promptOptions(b)
	}
}
func main() {
	var myBill Bill
	myBill = createBill()
	promptOptions(myBill)
	// fmt.Println(myBill.format())
}
