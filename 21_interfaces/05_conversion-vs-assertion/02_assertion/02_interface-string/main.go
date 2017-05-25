package main

import "fmt"

func main()  {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Println(str)
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("%T\n", str)
		fmt.Printf("value is not a string\n")
	}
}