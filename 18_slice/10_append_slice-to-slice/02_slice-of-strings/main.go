package main

import "fmt"

func main()  {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	mySlice = append(mySlice, "Saturday")

	fmt.Println(mySlice)
}