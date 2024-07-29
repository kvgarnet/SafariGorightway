package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func mightBreak() {
	if rand.Float32() > 0.5 {
		panic(errors.New("That broke"))
	}
	fmt.Println("succeeded")
}

func tryThat() {
	defer func() {
		fmt.Println("In deferred function")
		e := recover()
		if e != nil {
			fmt.Println("recovering from", e)
		} else {
			fmt.Println("no problem")
		}
	}()
	fmt.Println("trying")
	mightBreak()
	fmt.Println("back again")
}

func main() {
	tryThat()
	fmt.Println("back in main")
}
