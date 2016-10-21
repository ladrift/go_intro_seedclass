package main

import "fmt"

func main() {
	// main START OMIT
	sum := 0
	gen := fibs()
	for fib := gen(); fib < 4e6; fib = gen() {
		if fib%2 == 0 {
			sum += fib
		}
	}
	fmt.Println(sum)
	// main END OMIT
}

// fibs START OMIT
// fibs return a closure which generate fibonacci-
// number at each invocation.
func fibs() func() int {
	prev := 0
	curr := 1
	return func() int {
		tmp := curr
		curr += prev
		prev = tmp
		return curr
	}
}

// fibs END OMIT
