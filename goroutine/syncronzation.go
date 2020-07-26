package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("new routine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	foo(&wg)
	wg.Add(1)
	foo(&wg)
	wg.Done()
	fmt.Println("Main resource")
}
