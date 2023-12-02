package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	names := []string{"Kunj", "Cipher_KT", "r00t3d_KT"}
	for i := 0; i < len(names); i++ {
		fmt.Println(i+1, "->", names[i])
	}
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}
	for index, _ := range names {
		fmt.Printf("the index is %v \n", index)
	}
}
