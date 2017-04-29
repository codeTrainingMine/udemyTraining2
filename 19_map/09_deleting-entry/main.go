package main

import "fmt"

func main()  {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Bueonos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting, "len", len(myGreeting))
	delete(myGreeting, 2)
	fmt.Println(myGreeting, "len", len(myGreeting))

	myGreeting2 := map[string]string{
		"Tim": "Good morning!",
		"Jenny": "Bonjour!",
		"Peter": "Bueonos dias!",
		"Howard": "Bongiorno!",
	}

	fmt.Println(myGreeting, "len", len(myGreeting2))
	delete(myGreeting2, "Jenny")
	fmt.Println(myGreeting, "len", len(myGreeting2))
	//fmt.Println(myGreeting, "len", len(myGreeting2), "cap", cap(myGreeting2)) // invalid argument myGreeting2 (type map[string]string) for cap


}