package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func names() (string, string) {
	return "jhon", "torres"
}

// call by value
func foo(y int) {
	y = y + 1
}

// cal by reference
func goo(y *int) {
	*y = *y + 1
}

func doneOperation(myFun func(int, int) int, a, b int) int {
	return myFun(a, b)
}

func main() {
	defer fmt.Println("bye bye")

	divide := func(x, y int) (int, error) {
		if y == 0 {
			return 0, errors.New("y cannot be zero")
		}
		return x / y, nil
	}

	a, _ := divide(5, 2)
	fmt.Println(a)
}
