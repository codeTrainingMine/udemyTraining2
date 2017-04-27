package main

import "fmt"

func main()  {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	// experiment
	student = append(student, "aaa")
	fmt.Println(student)
}