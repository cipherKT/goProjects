package main

import "fmt"

func main() {
	name := "Cipher"
	age := 19
	fmt.Println("My name is ", name, " and, I am ", age, " years old!!")
	fmt.Printf("My name is %v and, I am %v years old!!\n", name, age) // %v -> format specifier for variable
	// %q -> format specifier for strings
	// %T -> format specifier for type
	// %f -> for float %0.1f -> float upto 1 decimal
	fmt.Printf("type of the variable age is %T\n", age)
	var str = fmt.Sprintf("My name is %v and, I am %v years old!!\n", name, age) // saves the formatted string in one variable
	fmt.Print("saved string is : ", str)
}
