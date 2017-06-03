package main

import (
	"os"
	"log"
	"fmt"
)

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("got to here")
}