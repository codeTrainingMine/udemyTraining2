package main

import "fmt"

func main()  {
	letter := rune("A"[0])
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	fmt.Println("============")
	letter2 := "A"[0]
	fmt.Println(letter2)
	fmt.Printf("%T \n", letter2)

}