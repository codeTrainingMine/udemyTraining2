package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape)  {
	fmt.Println("area", s.area())
}

func main()  {
	c := circle{5}
	info(&c)
	var c2 *circle
	c2 = &circle{6}
	info(c2)
}