package main

import "fmt"

func main() {
	println("arrays declarations")
	println("make")
	primes := []int{2, 3, 5, 7, 11}
	for index, prime := range primes {
		fmt.Printf("%d %d\n", index, prime)
	}
	var fibo [5]int
	fibo[0] = 0
	fibo[1] = 1
	fibo[2] = 1
	fibo[3] = 2
	fibo[4] = 4
	var facto = make([]int, 0, 5)
	facto = append(facto, 1)
	facto = append(facto, 1)
	facto = append(facto, 2)
	facto = append(facto, 6)
	facto = append(facto, 24)
	facto = append(facto, 120)

	println("slices")
	println("A 'windows' on an underlying array")
	s1 := primes[1:3]
	s2 := primes[2:5]
	println("s1", len(s1), cap(s1))
	println("s2", len(s2), cap(s2))

	println("test")
	zeros := make([]int, 10)
	zeros = append(zeros, 2)
	println("zeros", len(zeros))

	var zerosA [5]int
	println("zerosA", len(zerosA))

	println("Hash table")
	println("----------")
	agesFamily := map[string]int{"jhon": 27, "karen": 22, "jesus": 60, "aura": 65}
	println(agesFamily)
	println(agesFamily["jhon"])
	delete(agesFamily, "karen")

	idMap2 := make(map[string]int)
	idMap2["idMap"] = 15
	asd, ok := idMap2["jhon"]
	if !ok {
		println("jhon isn't in idMap", asd)
	}

	// struct
	type Person struct {
		Name    string
		Address string
		phone   int
	}
	var p1 Person
	p1.Name = "joe"

}
