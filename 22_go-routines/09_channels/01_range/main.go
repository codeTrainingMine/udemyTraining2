package main

import "fmt"

func main()  {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		//c <- 1  // panic send on closed channel
	}()

	for n := range c {
		fmt.Println(n)
	}
	//fmt.Println(<- c) // gives zero
	//fmt.Println(<- c) // gives zero
}