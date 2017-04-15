package main

import
(
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", page)
	res.Body.Close()
}