package main

import "fmt"

var x int

func earth()  {
	x++
	fmt.Println("4", x)
}

func main()  {
	fmt.Println("6", x)
	{
		defer earth()
	}
	fmt.Println("7", x)
}