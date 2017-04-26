package main

import "fmt"

var x int

func hello()  {
	fmt.Println("1", x)
	world()
	fmt.Println("5", x)
}

func world()  {
	fmt.Println("2", x)
	defer earth()
	fmt.Println("3", x)
}

func earth()  {
	x++
	fmt.Println("4", x)
}

func main()  {
	fmt.Println("6", x)
	hello()
	fmt.Println("7", x)
}