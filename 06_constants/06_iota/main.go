package main

import "fmt"

const (
	_ = iota          // 0
	b = iota * 10     // 1 * 10
	x string = "aa"
	c = iota * 10 	  // 2 * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
}