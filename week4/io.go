package main

import "os"

func main() {

	// read
	f, err := os.Open("dt.txt")
	if err != nil {

	}
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	println(nb)
	f.Close()

	// create
	f, err = os.Create("outfile.txt")
	bar := []byte{1, 2, 3}
	nb, err = f.Write(bar)
	nb, err = f.WriteString("h1")
	f.Close()
}
