package main

import "fmt"

func main()  {
	// slice
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	// array
	myArray := [3]int{1, 3, 5}
	fmt.Printf("%T\n", myArray)
	fmt.Println(myArray)
}