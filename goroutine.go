package main

import "fmt"

func main() {
	// main START OMIT
	var a string

	go func() {
		a = "hello, world"
		fmt.Println("sub routine exit")
	}()
	fmt.Println("main routine exit")
	fmt.Println(a)
	// main END OMIT
}
