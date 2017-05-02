package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main()  {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	if p1 == p2 {
		fmt.Println("equal")
	}
	p3 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p3.first, p3.last, p3.age)
	if p2 == p3 {
		fmt.Println("equal #2")
	}

}

