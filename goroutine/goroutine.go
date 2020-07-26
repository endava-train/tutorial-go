package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("new goroutine")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("hello world!!!")
}
