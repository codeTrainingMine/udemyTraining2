package main

import (
	"fmt"
)

func main() {
	counter := 0;
	for i := 0; i < 100000; i++ {
		for j := 3; j <= 13; j++ {
			n := fact(j);
			if counter % 10000 == 0 {
				fmt.Println(n)
			}
			counter++
		}
	}
	fmt.Println("counter", counter)
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
