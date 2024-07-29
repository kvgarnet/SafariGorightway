package main

import "fmt"

// no default argument values, no named parameter passing
func add(a ...int) int { // shows up as a slice of int
	//func add(a, b int, m string, c int) int {
	//func add(a int, b int) int {
	//	return a + b
	var sum int
	for _, v := range a {
		sum += v
	}
	//return a + b
	return sum
}

func addAndCount(a ...int) (sum int, count int) {
	//var sum int
	count = len(a)
	for _, v := range a {
		sum += v
	}
	//return sum, len(a)
	//return sum, count
	return
}

func main() {
	//fmt.Printf("sum of 1 and 2 is %d\n", add(1, 2, "", 99))
	//fmt.Printf("sum of 1 and 2 is %d\n", add(1, 2))
	fmt.Printf("sum of 1 and 2 is %d\n", add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("sum of 1 to 10 is %d\n", add(nums...))
	sum, count := addAndCount(nums...)
	fmt.Printf("sum of nums is %d, count is %d\n", sum, count)

}
