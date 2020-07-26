package main

import (
	"strconv"
	"strings"
)

func main() {
	println("string")
	println("strings are immutable")
	println("a string is compose by rune")
	println("%s in order to print with format")
	x := "hello world!!!"
	println(x)
	println(strings.ToLower(x))
	println(strings.Compare("jhon", "torres"))
	println(strings.HasPrefix("jhon emmanuel torres toloza", "jhon"))
	println(strings.Index("jhon emmanuel torres toloza", "emmanuel"))
	println(strings.TrimSpace(" jhon jhon  "))
	theIntConverted, _ := strconv.Atoi("123")
	println(theIntConverted)
	theStringConverted := strconv.Itoa(123)
	println(theStringConverted)

	println("integer")
	println("int8, int16, int32, int64, uint8, uint16, uint32, uint64")

	println("type conventions")
	println("Convert type with T() operation")
	println("example x = int32(y)")

	println("float32 6 digits of precision")
	println("float64 15 digits of precision")

	println("complex number")
	myComplex := complex(2, 3)
	println(myComplex)
}
