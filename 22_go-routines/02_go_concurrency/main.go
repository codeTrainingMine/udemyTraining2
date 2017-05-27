package main

import (
	"fmt"
	"time"
)

func main()  {
	go foo()
	go bar()
	time.Sleep(1000)
}

func foo()  {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar()  {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}