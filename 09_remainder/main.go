package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)

	//x = 13 / 3
	//fmt.Println(x)

	var y float64 = 13.0 / 3.0
	fmt.Println(y)

	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 0; i <= 100; i++ {
		if i % 2 == 1 {
			fmt.Printf("%d is odd\n", i)
		} else {
			fmt.Printf("%d is even\n", i)
		}
	}
}