package main

import (
	"encoding/json"
	"io/ioutil"
)

type PersonJSON struct {
	LastName  string `json:"lname"`
	FirstName string `json:"fname"`
}

func main() {
	bytesdt, err := ioutil.ReadFile("week4/dt.txt")
	if err != nil {
		println(err)
		return
	}
	p1 := new(PersonJSON)
	err = json.Unmarshal(bytesdt, &p1)
	if err != nil {
		println(err)
		return
	}
	println(p1.LastName)
	println(p1.FirstName)
	println(string(bytesdt))
}
