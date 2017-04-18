package main

import "fmt"

func zero(z int) {
	z = 0
	fmt.Printf("z address %p\n", &z)
	fmt.Printf("z %d\n", z)
}

func main() {
	x := 5
	zero(x)
	fmt.Printf("x address %p\n", &x)
	fmt.Printf("x %d\n", x)
}