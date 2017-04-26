package main

import "fmt"

func main()  {
	num, even := half(1)
	fmt.Println("num", num, "even", even)
	num, even = half(2)
	fmt.Println("num", num, "even", even)
	num, even = half(3)
	fmt.Println("num", num, "even", even)
	num, even = half(4)
	fmt.Println("num", num, "even", even)
}

func half(n int) (float64, bool) {
	return float64(n / 2), n % 2 == 0
}