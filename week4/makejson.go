package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var tmpVariable string
	person := map[string]string{}
	fmt.Scan(&tmpVariable)
	person["name"] = tmpVariable
	fmt.Scan(&tmpVariable)
	person["address"] = tmpVariable
	jsonP1, err := json.Marshal(person)
	if err != nil {
		println(err)
		return
	}
	println(string(jsonP1))
}
