package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	f := factorial(in)

	counter := 0;
	for n := range f {
		//fmt.Println("***************** println loop")
		if counter % 10000 == 0 {
			fmt.Println(n)
		}
		counter++
		//fmt.Println("***************** println loop end")
	}
	fmt.Println("counter", counter)
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for j := 3; j <= 13; j++ {
				//fmt.Println("***************** gen loop")
				out <- j
				//fmt.Println("***************** gen loop end")
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	results := make([]chan int, 1000)
	var wg sync.WaitGroup

	counter := 0
	for n := range in {
		result := make(chan int)
		results = append(results, result)
		//fmt.Println("***************** outside fact loop", &result)
		go func(num int, res chan int) {
			//fmt.Println("***************** fact loop", &res)
			res <- fact(num)
			//fmt.Println("***************** fact loop end")
		}(n, result)
		counter++
	}
	//fmt.Println("counter", counter)
	wg.Add(counter)

	for _, ch := range results {
		//fmt.Println("***************** outside slice loop", &ch)
		go func(c chan int) {
			//fmt.Println("***************** slice loop", &c)
			out <- <-c
			wg.Done()
			//fmt.Println("***************** slice loop end")
		}(ch)
	}

	go func() {
		//fmt.Println("***************** wait loop")
		wg.Wait()
		close(out)
		//fmt.Println("***************** wait loop end")
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
