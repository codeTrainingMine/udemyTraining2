package main

import "fmt"

func main()  {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
	}()

	go func() {
		<- c
		<- c
	}()

	fmt.Printf("c %v\n", c)
	fmt.Println(&c)
}