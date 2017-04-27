package main

import "fmt"

func main()  {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(len(student))
	fmt.Println(cap(student))

	student = append(student, "121")
	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println(cap(student))
}