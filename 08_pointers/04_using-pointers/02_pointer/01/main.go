package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Printf("x %d\n", x)
}