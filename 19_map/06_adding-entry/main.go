package main

import "fmt"

func main()  {

	myGreeting := map[string]string{
		"Tim":		"a",
		"Jenny":	"b",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
}