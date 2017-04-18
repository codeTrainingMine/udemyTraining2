package main

import "fmt"

func zero(z int) {
	fmt.Printf("z %d\n", z)
	z = 0
	fmt.Printf("z %d\n", z)
}

func main() {
	x := 5
	zero(x)
	fmt.Printf("x %d\n", x) // x is still 5
}