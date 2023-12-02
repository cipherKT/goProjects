package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Enter a number : ")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("The number provided by you is even")
	} else {
		fmt.Println("The number provided by you is odd")
	}
}
