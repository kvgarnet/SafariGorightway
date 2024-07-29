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
	// PROBLEM #2
	// I did not send the message to the control channel to tell the
	// main process that we're done!!! The following does that...
	// again, I think someone tried to tell me that, but I didn't
	// understand the help.
	control <- true
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
