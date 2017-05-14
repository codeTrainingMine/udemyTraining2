package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}

func main()  {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)

	r := sort.Reverse(studyGroup)
	sort.Sort(r)
	fmt.Println(studyGroup)
}