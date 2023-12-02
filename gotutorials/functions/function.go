package main

import (
	"fmt"
	"math"
)

func greetings(n string) {
	fmt.Printf("Hello World %v\n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func area(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// greetings("kunj")
	cycleNames([]string{"Kunj", "Cipher", "r00t3d"}, greetings)
	fmt.Println(area(12))
}
