package main

import (
	"fmt"
	"sort"
)

type people []string

func main()  {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}