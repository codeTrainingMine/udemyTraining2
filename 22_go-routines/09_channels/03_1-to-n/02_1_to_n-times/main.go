package main

import "fmt"

func main()  {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func(funcId int) {
			counter := 0
			for n := range c {
				fmt.Println("func", funcId, n)
				counter++
			}
			fmt.Println("************ counter", counter)
			done <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<-done
	}
}