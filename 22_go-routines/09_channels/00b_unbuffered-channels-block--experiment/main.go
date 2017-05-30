package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	c := make(chan int)

	done := false

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sender", i)
			c <- i
		}
		done = true
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("receiver", <-c)
		}
	}()

	counter := 0
	for ; !done ; {
		time.Sleep(time.Second)
		runtime.Gosched()
		counter++
	}
	fmt.Println("counter", counter)
}