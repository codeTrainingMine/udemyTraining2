package main

import "fmt"

func main() {
	i := 0
	for {
		if i % 100000000 == 0 {
			fmt.Println(i)
		}
		i++
	}
}