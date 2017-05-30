package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sender", i)
			c <- i
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("receiver", <-c)
		}
	}()

	time.Sleep(30 * time.Second)
}