package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name  string
	Phone string
	Age   string
}

func main() {
	p1 := Person{
		Name:  "jhon",
		Phone: "phone",
		Age:   "123",
	}
	fmt.Println(p1)

	barr, err := json.Marshal(p1)
	if err != nil {

	}

	var p2 Person
	err = json.Unmarshal(barr, &barr)
	if err != nil {

	}

	ioutil.ReadFile("text.txt")
	ioutil.WriteFile("outfile.txt", nil, 0777)

}
