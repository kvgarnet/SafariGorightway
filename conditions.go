package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 100
	fmt.Println("number is ", x)
	if x := rand.Intn(4); x > 2 { // local x must use := form
		fmt.Println("Big num!", x)
	} else {
		fmt.Println("not so big", x)
	}
	fmt.Println("x is ", x)

	switch y := rand.Intn(4); y {
	case 0:
		fmt.Println("It's zero")
	case 1, 2:
		fmt.Println("It's one or two")
		//fallthrough
	default:
		fmt.Println("not 0, 1, or 2")
	}
	//fmt.Println(y) // NOPE, y is local with :=

	// (x > 2) ? "Big" : "Small" NOPE, not allowed
	switch 0 {
	case rand.Intn(2):
		fmt.Println("Hit the first case")
	case rand.Intn(2):
		fmt.Println("Hit the second case")
	}

	fmt.Println("----------------------")

	switch { // with no expression here at all, equivalent to switch true. effectively a multi-way if/else
	case rand.Intn(2) > 0:
		fmt.Println("Hit the first case")
	case rand.Intn(2) > 0:
		fmt.Println("Hit the second case")
	default:
		fmt.Println("didn't hit either")
	}
}
