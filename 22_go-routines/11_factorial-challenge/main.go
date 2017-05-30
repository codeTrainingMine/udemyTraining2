package main

import "fmt"

func main()  {
	n := 4
	c := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()

	for j := range c {
		fmt.Println("Total:", j)
	}
}
