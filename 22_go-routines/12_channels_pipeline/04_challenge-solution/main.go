package main

import "fmt"

func main()  {
	var count uint64 = 100
	c := gen(count)
	fact := factorial(c, count)

	for n := range fact {
		fmt.Println(n)
	}
}

func gen(n uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		var i uint64
		for i = 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan uint64, count uint64) chan uint64  {
	out := make(chan uint64)
	fact := make(chan uint64)

	go func() {
		for n := range in {
			go func(num uint64) {
				var total uint64 = 1
				var i uint64
				for i = num; i > 0; i-- {
					total *= i
				}
				fact <- total
			}(n)
		}

		var i uint64
		for i = 1; i <= count; i++ {
			out <- (<-fact)
		}
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurreny?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
 */

