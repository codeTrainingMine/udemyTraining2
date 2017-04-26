package main

import "fmt"

var x int

func hello()  {
	fmt.Println("hello", x)
}

func world()  {
	fmt.Println("world", x)
}

func main()  {
	defer x++
	world()
	hello()
}