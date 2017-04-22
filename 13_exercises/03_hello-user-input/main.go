package main

import "fmt"

func main() {
	fmt.Print("enter your name:")
	var username string
	fmt.Scan(&username)
	fmt.Println("Hello " + username)
}