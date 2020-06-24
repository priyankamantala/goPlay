package main

import "fmt"

// create your own type and underlying type as int
type no int

var x no

func main() {
	fmt.Println("x =", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("x =", x)
}
