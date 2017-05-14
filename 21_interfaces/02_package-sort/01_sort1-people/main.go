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

	//fmt.Println(studyGroup.Len())
	//fmt.Println(studyGroup.Less(1, 2))
	//fmt.Println(studyGroup.Less(2, 1))
	//fmt.Println(studyGroup.Less(2, 2))
	//studyGroup.Swap(1, 2)

	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}