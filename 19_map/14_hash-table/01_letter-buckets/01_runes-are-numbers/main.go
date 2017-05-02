package main

import "fmt"

func main()  {
	letter := 'A'
	fmt.Println(letter)
	fmt.Println(string(letter))
	fmt.Printf("%T \n", letter)

	fmt.Println("====")
	letter2 := "A"
	fmt.Println(letter2)
	fmt.Println(letter2[0])
	fmt.Printf("%T \n", letter2)

	fmt.Println("====")
	letter3 := "ABCDEFGHIJKL"
	fmt.Println(letter3)
	fmt.Println(letter3[0:3])
	fmt.Println(letter3[3:6])
	fmt.Println(letter3[:6])
	fmt.Println(letter3[3:])
	fmt.Println(letter3[:])
	letter4 := letter3[:]
	fmt.Printf("%T \n", letter4)

}