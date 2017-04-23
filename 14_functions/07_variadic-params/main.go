package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for i, v := range sf {
		fmt.Println("iterating in range loop, v is", v, "i is", i)
		total += v
	}
	return total / float64(len(sf))
}