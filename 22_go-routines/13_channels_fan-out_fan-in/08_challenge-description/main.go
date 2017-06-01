package main

import (
	"fmt"
	"sync"
)

func main()  {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in
	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	for n := range merge(xc...) {
		fmt.Println("******** merge loop")
		fmt.Println(n)
		fmt.Println("******** merge loop end")
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		//for i := 0; i < 10; i++ {
			for j := 3; j < 5; j++ {
				fmt.Println("******** gen loop")
				out <- j
				fmt.Println("******** gen loop end")
			}
		//}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	xc := make([]<-chan int, 0)
	for i := 0; i < n; i++ {
		fmt.Println("******** append loop")
		xc = append(xc, factorial(in))
		fmt.Println("******** append loop end")
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("******** fact loop")
			out <- fact(n)
			fmt.Println("******** fact loop end")
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int  {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			fmt.Println("******** func loop")
			out <- n
			fmt.Println("******** func loop end")
		}
		wg.Done()
	}

	fmt.Println("len(cs)", len(cs))
	wg.Add(len(cs))
	for _, c := range cs {
		fmt.Println("******** output loop")
		go output(c)
		fmt.Println("******** output loop end")
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		fmt.Println("******** wait")
		wg.Wait()
		close(out)
		fmt.Println("******** wait end")
	}()
	return out
}