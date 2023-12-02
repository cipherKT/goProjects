package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello , my friends how are you!"
	fmt.Println(strings.Contains(greetings, "my friends"))
	fmt.Println(strings.ReplaceAll(greetings, "e", "3"))
	ages := []int{12, 32, 13, 43, 34, 56}
	sort.Ints(ages)
	fmt.Println(ages)
}
