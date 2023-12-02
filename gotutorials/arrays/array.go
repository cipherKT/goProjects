package main

import "fmt"

func main() {
	//arrays :
	// var ages [3]int = [3]int{20,25,30}
	var ages = [3]int{20, 25, 30} //shorthand declaration of array in GO
	name := [3]string{"Kunj", "Cipher", "r00t3d"}
	fmt.Println(ages, "is an array of type int and length", len(ages))
	fmt.Println(name, "is an array of type string and length", len(name))

	//slices (use under the array hood)
	var scores = []int{100, 50, 85, 70}
	fmt.Println("before appending", scores)
	scores = append(scores, 200) // this returns new slice of type int so we need to store it in a variable
	fmt.Println("after appending", scores)
	fmt.Printf("%T \n", scores)

	//slice ranges
	rangeOne := name[1:3] // this includes the element at position 1 and ignores the element at position 3
	fmt.Println(rangeOne)
}
