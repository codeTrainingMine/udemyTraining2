package main

import "fmt"

func main()  {
	fmt.Println(greatest(34,22,54,76,56,44,22,4343,34,22,1,55))
}

func greatest(nums ...int) int {
	var found int
	for _, num := range nums {
		//fmt.Println(num)
		if num > found {
			found = num
		}
	}
	return found;
}