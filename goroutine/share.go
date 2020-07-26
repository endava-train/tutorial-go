package main

import (
	"fmt"
	"sync"
)

var i = 0
var wg sync.WaitGroup
var mut sync.Mutex

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg.Add(1)
	go inc()
	wg.Add(1)
	go inc()
	wg.Wait()
	fmt.Println(i)
}
