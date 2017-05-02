package main

import "fmt"

func main()  {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)

	fmt.Println("======")
	letter2 := rune("Hello"[0])
	fmt.Println(letter2)
}