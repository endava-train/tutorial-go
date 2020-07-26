package main

import (
	"fmt"
	"runtime"
	"time"
)

func handle(queue chan int, workerId int) {
	for r := range queue {
		time.Sleep(time.Second)
		fmt.Println(r, workerId)
	}
}

func Serve(clientRequests chan int) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go handle(clientRequests, i)
	}
}

func main() {
	clientRequests := make(chan int)
	Serve(clientRequests)
	for i := 0; i < 100; i++ {
		clientRequests <- i * 20
	}
}
