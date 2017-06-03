package main

import (
	"os"
	"fmt"
	//"log"
)

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		//log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.