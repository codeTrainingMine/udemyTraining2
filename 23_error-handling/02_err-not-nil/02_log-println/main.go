package main

import (
	"os"
//"fmt"
//	"log"
	"log"
	"fmt"
)

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	fmt.Println("got to here")
}

// Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ... the Fatal functions call os.Exit(1) after writing the log message ... the Panic functions call panic after writing the log message.