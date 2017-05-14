package main

import (
	"fmt"
	"sort"
)

type people []string

func main()  {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	r := sort.Reverse(sort.StringSlice(s))
	sort.Sort(r)
	fmt.Println(s)
}