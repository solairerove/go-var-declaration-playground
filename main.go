package main

import "fmt"

var y = 42
var z int

func main() {

	x := 23

	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println(y)
}
