package main

import "fmt"

func main()  {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x))
	fmt.Println(int(y) + x)
	//fmt.Println(y + x) // error mismatched types float64 and int
	// conversion: int to float64
}