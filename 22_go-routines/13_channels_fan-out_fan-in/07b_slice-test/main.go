package main

import "fmt"

func main()  {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)

	for _, v := range s {
		fmt.Println(v)
	}
}