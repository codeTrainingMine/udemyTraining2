package main

import (
	"fmt"
	"sync"
)

func main()  {
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		//for i := 0; i < 10; i++ {
			for j := 1; j <= 3; j++ {
				out <- j
			}
		//}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int  {
	out := make(chan int)
	results := make([]chan int, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for n := range in {
			result := make(chan int)
			results = append(results, result)
			go func(num int, res chan int) {
				fmt.Println("***************** fact loop")
				res <- fact(num)
			}(n, result)
		}
	}()

	for _, ch := range results {
		go func(c chan int) {
			fmt.Println("***************** slice loop")
			out <- <- c
			wg.Done()
		}(ch)
	}

	go func() {
		fmt.Println("***************** wait loop")
		wg.Wait()
		close(out)
	}()

	return out
}

func fact(n int) int  {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
