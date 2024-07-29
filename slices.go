package main

import "fmt"

func main() {
	// allocate an array
	numbers := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("numbers is %#v\n", numbers)
	// take a slice
	//fromOne := numbers[1:3]
	//fromOne := []int{1, 2}
	fromOne := make([]int, 2, 4)
	fromOne[0] = 1
	fromOne[1] = 2
	fmt.Printf("fromOne is %#v, type is %T\n", fromOne, fromOne)
	fromOne[0] = 99
	fmt.Printf("numbers is %#v\n", numbers)
	fmt.Printf("fromOne is %#v, type is %T, len is %d, cap is %d\n",
		fromOne, fromOne, len(fromOne), cap(fromOne))
	fromOne = append(fromOne, 101, 102, 103)
	fmt.Println("numbers is %#v\n", numbers)
	fmt.Printf("fromOne is %#v, type is %T, len is %d, cap is %d\n",
		fromOne, fromOne, len(fromOne), cap(fromOne))
}
