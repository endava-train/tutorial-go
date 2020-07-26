package main

import "fmt"

var x = 4

func f() {
	fmt.Printf("%d\n", x)
}

func g() {
	fmt.Printf("%d\n", x)
}

func a() {
	x := 5
	prime := 2
	fmt.Println(x)
	fmt.Println(prime)
}

func b() {
	println("x = global scope", x)
	println("I cannot access to prime")
}

func main() {
	f()
	g()
	a()
	b()
}
