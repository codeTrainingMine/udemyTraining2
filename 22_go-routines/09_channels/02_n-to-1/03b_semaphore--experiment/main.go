package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int)
	done1, done2 := false, false

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done1 = true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done2 = true
	}()

	go func() {
		for ;!done1 && !done2; {
			time.Sleep(time.Second)
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}