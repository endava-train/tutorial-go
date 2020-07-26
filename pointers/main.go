package main

import "fmt"

type Point struct {
	x int
	y int
}

// inline comment
/**
block comment
*/

func main() {
	fmt.Println("Pointers")
	fmt.Println("--------")
	fmt.Println("A pointer is an address to data in memory")
	fmt.Println("& operator returns the address of a variable / function")
	fmt.Println("* operator reutrns data at an address (deferencing)")

	var x int = 1
	var y int
	var ip *int

	ip = &x // ip now points to x
	y = *ip // y is now 1. this do a copy

	fmt.Println("x =", x, &x)
	fmt.Println("y =", y, &y)
	fmt.Println("ip =", ip, &ip, *ip)
	fmt.Println()

	fmt.Println("New operator")
	fmt.Println("new() function creates a variable and returns a pointer to the variable")

	ptr := new(int)
	fmt.Println("ptr", ptr)
	*ptr = 3
	fmt.Println("ptr", ptr)

	zeroPoint := new(Point)
	println(zeroPoint.x, zeroPoint.y)

	identity := Point{
		x: 1,
		y: 1,
	}

	println(identity.x, identity.y)

}
