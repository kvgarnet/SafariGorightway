package main

import "fmt"

func main() {
	a := 10
	//++a // NOPE!
	a++ // NOT and EXPRESSION
	//b := a++ // NOPE!!!
	// also assignments are not expressions
	//b := (a = 99) // ALSO NOT ALLOWED
	a += 10 // allowed, but not an expression
	fmt.Println(a)

	for x := 0; x < 3; x++ {
		fmt.Println("x is ", x)
	}

	var x int = 3
	for ; x > 0; x-- {
		fmt.Println("x is now ", x)
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//for i, v := range nums {
	for _, v := range nums { // underscore as placeholder for value to be assigned but not needed!
		//fmt.Printf("index %d, value %d\n", i, v)
		fmt.Printf("value %d\n", v)
	}

	var y, z int
	y, z = 99, 100
	fmt.Printf("x is %d, y is %d\n", y, z)
}
