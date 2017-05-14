package main

import (
	"fmt"
	//"sort"
	"sort"
)

type people []string

func main()  {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)

	sort.Strings([]string(studyGroup))
	fmt.Println(studyGroup)
}