package main

import (
	"fmt"
)


func main()  {
	count := make(chan int)
	done := make(chan bool)
	go incrementor("1", count, done)
	go incrementor("2", count, done)

	go closer(count, done)

	var finalCount int
	for range count {
		finalCount++
	}

	fmt.Println("Final Counter:", finalCount)
}

func closer(count chan int, done chan bool)  {
	<-done
	<-done
	close(count)
}

func incrementor(s string, count chan int, done chan bool) {
	for i := 0; i < 20; i++ {
		count <- i
		fmt.Println("Process: " + s + " printing:", i)
	}
	done <- true
}