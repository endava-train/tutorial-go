package main

import "fmt"

func main() {
	const x = 5
	if x == 5 {
		println("I am five")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("hello world %d\n", i)
	}

	// for like while
	i := 0
	for i < 5 {
		fmt.Printf("hello world %d\n", i)
		i++
	}

	for {
		fmt.Println("infinity look")
		break
	}

	switch x {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	default:
		println("another number")
	}

	switch {
	case x > 1:
		println("positive")
	case x < -1:
		println("negative")
	default:
		println("ciro")
	}

	println("scan reads user input")
	// Declaring some variables
	var name string
	var alphabet_count int

	// Calling Scan() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scan(&name)
	fmt.Scan(&alphabet_count)

	// Printing the given texts
	fmt.Printf("The word %s containg %d number of alphabets.",
		name, alphabet_count)
}
