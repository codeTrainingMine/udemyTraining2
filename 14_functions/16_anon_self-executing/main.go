package main

import "fmt"

func main()  {
	i := 12
	func(i int) {
		fmt.Println("I'm driving!", i)
	}(i)
}