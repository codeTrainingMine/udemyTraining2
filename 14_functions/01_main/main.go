package main

import (
	"fmt"
	"os"
)

// cant have arguments or return values
//func main(a int) {
//	fmt.Println("Hello world and your number: ", a)
//}

func main() {
	fmt.Println("Hello world!")
	a := os.Args[0]
	fmt.Println("os.Args[0]", a)
	b := os.Args[1]
	fmt.Println("os.Args[1]", b)
}


// main is the entry point to your program