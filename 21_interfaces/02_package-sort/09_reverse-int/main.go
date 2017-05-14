package main

import (
	"fmt"
	"sort"
)

type people []string

func main()  {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)

	r := sort.Reverse(sort.IntSlice(n))
	sort.Sort(r)
	fmt.Println(n)
}