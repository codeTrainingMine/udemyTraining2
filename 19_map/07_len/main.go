package main

import "fmt"

func main()  {

	myGreeting := map[string]string{
		"Tim":		"a",
		"Jenny":	"b",
	}

	fmt.Println(len(myGreeting))
	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))
}