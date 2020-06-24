package main

import "fmt"

// package level declaration using VAR keyword

var x int
var y string
var z bool

func main() {
	// zero values will be assigned by the compiler
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
