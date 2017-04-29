package main

import "fmt"

func main()  {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Bueonos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val:", val)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	if val, exists := myGreeting[5]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val:", val)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}


}