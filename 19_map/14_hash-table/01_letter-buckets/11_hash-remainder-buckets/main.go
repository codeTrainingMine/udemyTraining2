package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 12)
	words := make([]string, 12)
	// Loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++

		if len(words[n]) < 50 {
			words[n] += " " + scanner.Text()
		}
	}
	fmt.Println(buckets)

	for i := 0; i <= 11; i++ {
		fmt.Println(i, words[i])
	}
}

func hashBucket(word string, buckets int) int  {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}