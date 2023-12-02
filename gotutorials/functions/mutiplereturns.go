package main

import (
	"fmt"
	"math"
)

func getCircumferenceAndArea(r float64) (float64, float64) {
	return 2 * math.Pi * r, math.Pi * r * r
}
func main() {
	r := 5.0
	cir, ar := getCircumferenceAndArea(r)
	fmt.Printf("Circumnference of the circle is %v \nArea of the circle is %v", cir, ar)

}
