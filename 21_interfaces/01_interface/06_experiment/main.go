package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main()  {
	var w io.Writer
	fmt.Printf("%T, %v \n", w, w)

	w = os.Stdout
	fmt.Printf("%T, %v \n", w, w)

	w = new(bytes.Buffer)
	fmt.Printf("%T, %v \n", w, w)

	w = nil
	fmt.Printf("%T, %v \n", w, w)
}