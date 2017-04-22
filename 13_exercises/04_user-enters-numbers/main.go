package main

import "fmt"

func main() {
	fmt.Print("enter small number:")
	var smallNum int
	fmt.Scan(&smallNum)

	fmt.Print("enter larger number:")
	var largerNum int
	fmt.Scan(&largerNum)

	fmt.Println("remainder of larger num divided by smaller num:")
	fmt.Println(largerNum % smallNum)


}