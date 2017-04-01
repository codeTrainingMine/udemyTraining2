package main

import (
	"fmt"
	"udemyTraining2/04_scope/01_package-scope/01/mypack"
)

var x = 42

func main() {
	fmt.Println(x)
	foo()
	mypack.Foo2()
	mypack.Foo3()
	//Food()
}

func foo() {
	fmt.Println(x)
	//Foo2()
}