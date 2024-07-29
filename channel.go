package main

import "fmt"

func producer(out chan int) {
	for i := 0; i < 100; i++ {
		out <- i + 10
	}
	out <- -1
	fmt.Println("producer shutting down")
}

func consumer(in chan int, control chan bool) {
	for {
		v := <-in
		if v < 0 {
			break
		}
		fmt.Println("value is ", v)
	}
	fmt.Println("consumer shutting down")
}

func main() {
	data := make(chan int)
	control := make(chan bool)

	go producer(data)
	go consumer(data, control)
	// wait for control message
	x := <-control
	fmt.Println(x)
	fmt.Println("Main shutting down")
}
