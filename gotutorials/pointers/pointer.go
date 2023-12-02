package main

import "fmt"

func updateName(n *string) {
	*n = "Cipher"
}
func main() {
	name := "r00t3d"
	namePointer := &name
	fmt.Println(name)
	updateName(namePointer)
	fmt.Println(name)

}
