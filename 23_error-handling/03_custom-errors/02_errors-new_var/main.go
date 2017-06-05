package main

import (
	"errors"
	"log"
	"fmt"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main()  {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error)  {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}