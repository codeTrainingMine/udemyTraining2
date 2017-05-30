package main

import "fmt"

func main()  {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a dealock.
// Can you determine why?
// And what would you do to fix it?