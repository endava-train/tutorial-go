package main

import "time"

func prod(v1, v2 int, c chan int) {
	time.Sleep(time.Second)
	c <- v1 * v2
}

func main() {
	println("Communication")
	println("-------------")
	println("Transfer data between goroutines")
	println("Channels are typed")
	println("Send and receive data using the <- operator")
	c := make(chan int)
	println("Send data on a channel")
	//c <- 3
	println("Receive data from a channel")
	//x := <- c
	//println("Data receive", x)

	println("This code do synchronization")
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <-c
	b := <-c
	println(a * b)

	println("Channel Capacity")
	println("Channels can contains a limited number of objects")
	println("  Default size 0")
	println("Capacity is the number of objects it can hold in transit")
	println("Optional arguments to make defines channel capacity")
	//c := make(chan int, 3)
	println("Sending only blocks if buffer is full")
	println("Receiving only blocks if buffer is empty")
}
