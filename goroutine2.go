package main

import "fmt"

func main() {
	// main START OMIT
	var a string

	c := make(chan int, 0) // unbuffered channel
	go func() {
		a = "hello, world"
		fmt.Println("sub routine exit")
		<-c
	}()
	c <- 0
	fmt.Println("main routine exit")
	fmt.Println(a)
	// main END OMIT
}
