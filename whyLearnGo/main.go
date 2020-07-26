package main

// this function must be documented
func Documented() {}

func main() {
	println("Advantages of go")
	println("----------------")
	println("1. Code run fast")
	println("2. Garbage collector")
	println("3. Simpler objects")
	println("4. Concurrency is efficient")
	println()

	println("Objects in Go")
	println("-------------")
	println("* Go does not use the term class")
	println("* Go uses **struct** with associated methods")
	println("* Simplified implementation of clases")
	println("  * No inheritance")
	println("  * No constructors")
	println("  * No generics")
	println()

	println("Concurrency in go")
	println("-----------------")
	println("Go includes concurrency primitives")
	println("Goroutines represents concurrent tasks")
	println("Channels are used to communicate between tasks")
	println("Select enables task synchronization")
	println("Concurrency primitives are efficiency and easy to use")
	println()

	println("Workspaces")
	println("----------")
	println("* src - Contains source code files")
	println("* pkg - Contains packages (libraries)")
	println("* bin - Contains binaries")
	println("GOPATH is the entorn variable ")
	println()

	println("Packages")
	println("--------")
	println("* group of related source files")
	println("* each package can be imported by other packages")
	println("* enables software reuse")
	println("* there must be one package called **main.go**")
	println()

	println("The go tool")
	println("-----------")
	println("go build ... -> creates an executable")
	println("go doc")
	println("go fmt")
	println("go get")
	println("go list -> installed packages")
	println("go run -> compiles .go files and runs the executable")

}
